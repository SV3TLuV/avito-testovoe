package app

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"log"
	"tender_api/src/internal/config"
	"tender_api/src/internal/db/postgres"
	"tender_api/src/internal/repository"
	"tender_api/src/internal/service"
)

type ServiceProvider struct {
	config *config.Config

	postgres  *pgxpool.Pool
	trManager *manager.Manager

	tenderRepo    repository.TenderRepository
	tenderService service.TenderService

	bidRepo    repository.BidRepository
	bidService service.BidService

	employeeRepo repository.EmployeeRepository
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (p *ServiceProvider) Config() *config.Config {
	if p.config == nil {
		cfg, err := config.FromEnv()
		if err != nil {
			log.Fatal(errors.Wrap(err, "init config"))
		}
		p.config = cfg
	}
	return p.config
}

func (p *ServiceProvider) Postgres() *pgxpool.Pool {
	if p.postgres == nil {
		db, err := postgres.NewDB(context.Background(), p.Config().PostgresConn)
		if err != nil {
			log.Fatal(errors.Wrap(err, "init postgresql pool"))
		}
		p.postgres = db
	}
	return p.postgres
}

func (p *ServiceProvider) TransactionManager() *manager.Manager {
	if p.trManager == nil {
		p.trManager = manager.Must(trmpgx.NewDefaultFactory(p.postgres))
	}
	return p.trManager
}

func (p *ServiceProvider) TenderService() service.TenderService {
	if p.tenderService == nil {
		p.tenderService = service.NewTenderService(
			p.TenderRepo(),
			p.EmployeeRepo(),
			p.TransactionManager())
	}
	return p.tenderService
}

func (p *ServiceProvider) TenderRepo() repository.TenderRepository {
	if p.tenderRepo == nil {
		p.tenderRepo = repository.NewTenderRepository(
			p.Postgres(),
			trmpgx.DefaultCtxGetter,
			p.TransactionManager())
	}
	return p.tenderRepo
}

func (p *ServiceProvider) BidService() service.BidService {
	if p.bidRepo == nil {
		p.bidService = service.NewBidService(
			p.BidRepo(),
			p.EmployeeRepo())
	}
	return p.bidService
}

func (p *ServiceProvider) BidRepo() repository.BidRepository {
	if p.bidRepo == nil {
		p.bidRepo = repository.NewBidRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.bidRepo
}

func (p *ServiceProvider) EmployeeRepo() repository.EmployeeRepository {
	if p.employeeRepo == nil {
		p.employeeRepo = repository.NewEmployeeRepository(
			p.Postgres(),
			trmpgx.DefaultCtxGetter)
	}
	return p.employeeRepo
}
