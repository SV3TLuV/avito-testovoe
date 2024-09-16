package repository

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
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

type bidRepository struct {
	pool   *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewBidRepository(pool *pgxpool.Pool, getter *trmpgx.CtxGetter) *bidRepository {
	return &bidRepository{
		pool:   pool,
		getter: getter,
	}
}

func (repo *bidRepository) GetMy(ctx context.Context, limit, offset uint, employeeId uuid.UUID) ([]model.Bid, error) {
	query := goqu.Dialect("postgres").
		From("bid").
		Where(goqu.Ex{"author_id": employeeId}).
		Order(goqu.I("name").Asc()).
		Limit(limit).
		Offset(offset)

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	bids := make([]model.Bid, 0)
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Select(ctx, tr, &bids, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return bids, nil
}

func (repo *bidRepository) GetTenderBidList(ctx context.Context, tenderID uuid.UUID, limit, offset uint,
	employeeID, organizationID uuid.UUID) ([]model.Bid, error) {
	query := goqu.Dialect("postgres").
		Select("bid.*").
		From("bid").
		Join(
			goqu.T("tender"),
			goqu.On(goqu.Ex{"tender.id": goqu.I("bid.tender_id")})).
		Where(
			goqu.And(
				goqu.Ex{"bid.tender_id": tenderID},
				goqu.Or(
					goqu.Ex{"bid.author_id": employeeID},
					goqu.Ex{"tender_id.organization_id": organizationID}))).
		Order(goqu.I("bid.name").Asc()).
		Limit(limit).
		Offset(offset)

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	bids := make([]model.Bid, 0)
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Select(ctx, tr, &bids, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return bids, nil
}

func (repo *bidRepository) GetStatus(ctx context.Context, bidID uuid.UUID) (*model.BidStatusResponse, error) {
	query := goqu.Dialect("postgres").
		Select("bid.status", "tender.organization_id", "bid.author_id").
		From("bid").
		Join(
			goqu.T("tender"),
			goqu.On(goqu.Ex{"tender.id": goqu.I("bid.tender_id")})).
		Where(goqu.Ex{"bid.id": bidID})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var status model.BidStatusResponse
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &status, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "bid not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &status, nil
}

func (repo *bidRepository) GetTenderReviews(ctx context.Context, limit, offset uint,
	tenderID, authorID, requesterOrganizationID uuid.UUID) ([]model.BidReview, error) {
	query := goqu.Dialect("postgres").
		Select("bid_review.*").
		From("bid").
		Join(
			goqu.T("bid_review"),
			goqu.On(goqu.Ex{"bid_review.bid_id": goqu.I("bid.id")})).
		Join(
			goqu.T("tender"),
			goqu.On(goqu.Ex{"tender.id": goqu.I("bid.tender_id")})).
		Where(
			goqu.And(
				goqu.Ex{"bid.tender_id": tenderID},
				goqu.Ex{"bid.author_id": authorID})).
		Where(goqu.Ex{"tender.organization_id": requesterOrganizationID}).
		Limit(limit).
		Offset(offset)

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	reviews := make([]model.BidReview, 0)
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Select(ctx, tr, &reviews, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return reviews, nil
}

func (repo *bidRepository) GetById(ctx context.Context, bidID uuid.UUID) (*model.Bid, error) {
	query := goqu.Dialect("postgres").
		From("bid").
		Where(goqu.Ex{"id": bidID})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var bid model.Bid
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &bid, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "bid not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &bid, nil
}

func (repo *bidRepository) GetTenderOrganizationId(ctx context.Context, bidID uuid.UUID) (uuid.UUID, error) {
	query := goqu.Dialect("postgres").
		Select("tender.organization_id").
		From("bid").
		Join(
			goqu.T("tender"),
			goqu.On(goqu.Ex{"tender.id": goqu.I("bid.tender_id")})).
		Where(goqu.Ex{"bid.id": bidID})

	sql, args, err := query.ToSQL()
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to build query")
	}

	var organizationID uuid.UUID
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &organizationID, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return uuid.Nil, errors.Wrap(model.ErrNotFound, "bid not found")
	}
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to execute query")
	}

	return organizationID, nil
}

func (repo *bidRepository) GetByVersion(ctx context.Context, bidID uuid.UUID, version uint64) (*model.Bid, error) {
	if bid, err := repo.GetById(ctx, bidID); err == nil && bid.Version == version {
		return bid, nil
	}

	query := goqu.Dialect("postgres").
		From("bid_history").
		Where(goqu.And(
			goqu.Ex{"bid_id": bidID},
			goqu.Ex{"version": version}))

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var history model.BidHistory
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &history, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "bid not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	bid := converter.ToBidFromBidHistory(history)
	return &bid, nil
}

func (repo *bidRepository) Create(ctx context.Context, entity model.Bid) (*model.Bid, error) {
	record := converter.ToBidRecordFromBid(entity)
	record["status"] = enum.BidCreated

	query := goqu.Dialect("postgres").
		Insert("bid").
		Rows(record).
		Returning("bid.*")

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var bid model.Bid
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &bid, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &bid, nil
}

func (repo *bidRepository) Edit(ctx context.Context, entity model.Bid) (*model.Bid, error) {
	current, err := repo.GetById(ctx, entity.ID)
	if err != nil {
		return nil, err
	}

	if err = repo.archive(ctx, *current); err != nil {
		return nil, errors.Wrap(err, "failed to archive bid")
	}

	entity.Version = current.Version + 1
	record := converter.ToBidRecordFromBid(entity)
	record["status"] = enum.BidCreated

	query := goqu.Dialect("postgres").
		Update("bid").
		Set(record).
		Where(goqu.Ex{"id": entity.ID}).
		Returning("bid.*")

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var bid model.Bid
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &bid, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &bid, nil
}

func (repo *bidRepository) UpdateStatus(ctx context.Context, bidID uuid.UUID, status enum.BidStatus) (*model.Bid, error) {
	query := goqu.Dialect("postgres").
		Update("bid").
		Set(goqu.Record{
			"status": status,
		}).
		Where(goqu.Ex{"id": bidID}).
		Returning("bid.*")

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var bid model.Bid
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &bid, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &bid, nil
}

func (repo *bidRepository) SubmitDecision(ctx context.Context, bidID uuid.UUID, decision enum.BidDecision) error {
	query := goqu.Dialect("postgres").
		Update("bid").
		Set(goqu.Record{
			"decision": decision,
		}).
		Where(goqu.Ex{"id": bidID})

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

func (repo *bidRepository) Feedback(ctx context.Context, bidID uuid.UUID, feedback string) error {
	query := goqu.Dialect("postgres").
		Insert("bid_review").
		Rows(goqu.Record{
			"bid_id":      bidID,
			"description": feedback,
		})

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

func (repo *bidRepository) archive(ctx context.Context, entity model.Bid) error {
	bidHistory := converter.ToBidHistoryFromBid(entity)
	record := converter.ToBidHistoryRecordFromBidHistory(bidHistory)
	query := goqu.Dialect("postgres").
		Insert("bid_history").
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
