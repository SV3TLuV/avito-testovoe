package repository

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"tender_api/src/internal/converter"
	"tender_api/src/internal/model"
	"tender_api/src/internal/model/enum"
)

var _ TenderRepository = (*tenderRepository)(nil)

type tenderRepository struct {
	pool      *pgxpool.Pool
	getter    *trmpgx.CtxGetter
	trManager *manager.Manager
}

func NewTenderRepository(pool *pgxpool.Pool,
	getter *trmpgx.CtxGetter,
	trManager *manager.Manager) *tenderRepository {
	return &tenderRepository{
		pool:      pool,
		getter:    getter,
		trManager: trManager,
	}
}

func (repo *tenderRepository) GetList(ctx context.Context, limit, offset uint, serviceType []enum.TenderServiceType) ([]model.Tender, error) {
	query := goqu.Dialect("postgres").
		From("tender").
		Where(goqu.Ex{"status": enum.TenderPublished}).
		Order(goqu.I("name").Asc()).
		Limit(limit).
		Offset(offset)

	if len(serviceType) > 0 {
		query = query.Where(goqu.Ex{"service_type": goqu.Op{"IN": serviceType}})
	}

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	tenders := make([]model.Tender, 0)
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Select(ctx, tr, &tenders, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return tenders, nil
}

func (repo *tenderRepository) GetMy(ctx context.Context, limit, offset uint, employeeID uuid.UUID) ([]model.Tender, error) {
	query := goqu.Dialect("postgres").
		From("tender").
		Select("tender.*").
		Join(
			goqu.T("employee_tender"),
			goqu.On(goqu.Ex{"tender.id": goqu.I("employee_tender.tender_id")})).
		Where(goqu.Ex{"employee_tender.employee_id": employeeID}).
		Order(goqu.I("tender.name").Asc()).
		Limit(limit).
		Offset(offset)

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	tenders := make([]model.Tender, 0)
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Select(ctx, tr, &tenders, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return tenders, nil
}

func (repo *tenderRepository) GetStatus(ctx context.Context, tenderID uuid.UUID) (*model.TenderStatusResponse, error) {
	query := goqu.Dialect("postgres").
		Select("tender.status", "tender.organization_id").
		From("tender").
		Where(goqu.Ex{"id": tenderID})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var status model.TenderStatusResponse
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &status, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "tender not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &status, nil
}

func (repo *tenderRepository) GetById(ctx context.Context, tenderID uuid.UUID) (*model.Tender, error) {
	query := goqu.Dialect("postgres").
		From("tender").
		Where(goqu.Ex{"id": tenderID})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var tender model.Tender
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &tender, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "tender not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &tender, nil
}

func (repo *tenderRepository) GetByVersion(ctx context.Context, tenderID uuid.UUID, version uint64) (*model.Tender, error) {
	if tender, err := repo.GetById(ctx, tenderID); err == nil && tender.Version == version {
		return tender, nil
	}

	query := goqu.Dialect("postgres").
		From("tender_history").
		Where(goqu.And(
			goqu.Ex{"id": tenderID},
			goqu.Ex{"version": version}))

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var history model.TenderHistory
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &history, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "tender not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	tender := converter.ToTenderFromTenderHistory(history)
	return &tender, nil
}

func (repo *tenderRepository) Create(ctx context.Context, entity model.Tender, employeeID uuid.UUID) (*model.Tender, error) {
	var tender model.Tender
	err := repo.trManager.Do(ctx, func(ctx context.Context) error {
		tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)

		record := converter.ToTenderRecordFromTender(entity)
		record["status"] = enum.TenderCreated

		query := goqu.Dialect("postgres").
			Insert("tender").
			Rows(record).
			Returning("tender.*")

		sql, args, err := query.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		err = pgxscan.Get(ctx, tr, &tender, sql, args...)
		if err != nil {
			return errors.Wrap(err, "failed to execute query")
		}

		query = goqu.Dialect("postgres").
			Insert("employee_tender").
			Rows(goqu.Record{
				"employee_id": employeeID,
				"tender_id":   tender.ID,
			})

		sql, args, err = query.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "failed to execute query")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "transaction failed")
	}

	return &tender, nil
}

func (repo *tenderRepository) Edit(ctx context.Context, entity model.Tender) (*model.Tender, error) {
	current, err := repo.GetById(ctx, entity.ID)
	if err != nil {
		return nil, err
	}

	var tender model.Tender
	err = repo.trManager.Do(ctx, func(ctx context.Context) error {
		tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
		if err = repo.archive(ctx, *current); err != nil {
			return errors.Wrap(err, "failed to archive tender")
		}

		entity.Version = current.Version + 1
		record := converter.ToTenderRecordFromTender(entity)
		record["updated_at"] = goqu.L("CURRENT_TIMESTAMP")

		query := goqu.Dialect("postgres").
			Update("tender").
			Set(record).
			Where(goqu.Ex{"id": entity.ID}).
			Returning("tender.*")

		sql, args, err := query.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		err = pgxscan.Get(ctx, tr, &tender, sql, args...)
		if err != nil {
			return errors.Wrap(err, "failed to execute query")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "transaction failed")
	}

	return &tender, nil
}

func (repo *tenderRepository) UpdateStatus(ctx context.Context, tenderID uuid.UUID, status enum.TenderStatus) (*model.Tender, error) {
	query := goqu.Dialect("postgres").
		Update("tender").
		Set(goqu.Record{
			"status": status,
		}).
		Where(goqu.Ex{"id": tenderID}).
		Returning("tender.*")

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var tender model.Tender
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &tender, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &tender, nil
}

func (repo *tenderRepository) archive(ctx context.Context, entity model.Tender) error {
	tenderHistory := converter.ToTenderHistoryFromTender(entity)
	record := converter.ToTenderHistoryRecordFromTenderHistory(tenderHistory)
	query := goqu.Dialect("postgres").
		Insert("tender_history").
		Rows(record)

	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "failed to execute query")
	}

	return nil
}
