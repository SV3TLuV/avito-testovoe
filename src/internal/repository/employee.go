package repository

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"tender_api/src/internal/model"
)

var _ EmployeeRepository = (*employeeRepository)(nil)

type employeeRepository struct {
	pool   *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewEmployeeRepository(pool *pgxpool.Pool,
	getter *trmpgx.CtxGetter) *employeeRepository {
	return &employeeRepository{
		pool:   pool,
		getter: getter,
	}
}

func (repo *employeeRepository) GetByUsername(ctx context.Context, username string) (*model.Employee, error) {
	query := goqu.Dialect("postgres").
		From("employee").
		Where(goqu.Ex{"username": username})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var employee model.Employee
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &employee, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "employee not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &employee, nil
}

func (repo *employeeRepository) GetUserOrganization(ctx context.Context, username string) (*model.Organization, error) {
	query := goqu.Dialect("postgres").
		Select("organization.*").
		From("organization").
		Join(
			goqu.T("organization_responsible"),
			goqu.On(goqu.Ex{"organization_responsible.organization_id": goqu.I("organization.organization_id")})).
		Join(
			goqu.T("employee"),
			goqu.On(goqu.Ex{"employee.employee_id": goqu.I("organization_responsible.user_id")})).
		Where(goqu.Ex{"employee.username": username})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var organization model.Organization
	tr := repo.getter.DefaultTrOrDB(ctx, repo.pool)
	err = pgxscan.Get(ctx, tr, &organization, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(model.ErrNotFound, "organization not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &organization, nil
}
