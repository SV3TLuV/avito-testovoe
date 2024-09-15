package service

import (
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"tender_api/src/internal/model"
	"tender_api/src/internal/model/enum"
	"tender_api/src/internal/repository"
)

var _ TenderService = (*tenderService)(nil)

type tenderService struct {
	repo         repository.TenderRepository
	employeeRepo repository.EmployeeRepository
	trManager    *manager.Manager
}

func NewTenderService(repo repository.TenderRepository,
	employeeRepo repository.EmployeeRepository,
	trManager *manager.Manager) *tenderService {
	return &tenderService{
		repo:         repo,
		employeeRepo: employeeRepo,
		trManager:    trManager,
	}
}

func (s *tenderService) GetList(ctx context.Context, limit, offset uint, serviceType []enum.TenderServiceType) ([]model.Tender, error) {
	return s.repo.GetList(ctx, limit, offset, serviceType)
}

func (s *tenderService) GetMy(ctx context.Context, limit, offset uint, username string) ([]model.Tender, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return s.repo.GetMy(ctx, limit, offset, employee.ID)
}

func (s *tenderService) GetStatus(ctx context.Context, tenderID uuid.UUID, username string) (enum.TenderStatus, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	statusResponse, err := s.repo.GetStatus(ctx, tenderID)
	if err != nil {
		return "", err
	}

	if statusResponse.Status == enum.TenderPublished {
		return statusResponse.Status, nil
	}

	organization, err := s.employeeRepo.GetUserOrganization(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return "", errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return "", err
	}

	if statusResponse.OrganizationID != organization.ID {
		return "", errors.Wrap(model.ErrForbidden, "access denied to tender")
	}

	return statusResponse.Status, nil
}

func (s *tenderService) Create(ctx context.Context, entity model.Tender, username string) (*model.Tender, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganization(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}
	if organization.ID != entity.OrganizationID {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	created, err := s.repo.Create(ctx, entity, employee.ID)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *tenderService) Edit(ctx context.Context, entity model.Tender, username string) (*model.Tender, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganization(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}

	tender, err := s.repo.GetStatus(ctx, entity.ID)
	if err != nil {
		return nil, err
	}

	if organization.ID != tender.OrganizationID {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	updated, err := s.repo.Edit(ctx, entity)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *tenderService) UpdateStatus(ctx context.Context, tenderID uuid.UUID,
	status enum.TenderStatus, username string) (*model.Tender, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganization(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}

	tender, err := s.repo.GetStatus(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	if organization.ID != tender.OrganizationID {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	updated, err := s.repo.UpdateStatus(ctx, tenderID, status)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *tenderService) Rollback(ctx context.Context, tenderID uuid.UUID,
	version uint64, username string) (*model.Tender, error) {
	tenderToRollback, err := s.repo.GetByVersion(ctx, tenderID, version)
	if err != nil {
		return nil, err
	}

	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganization(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}
	if organization.ID != tenderToRollback.OrganizationID {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	updated, err := s.repo.Edit(ctx, *tenderToRollback)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
