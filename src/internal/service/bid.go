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

var _ BidService = (*bidService)(nil)

type bidService struct {
	bidRepo      repository.BidRepository
	tenderRepo   repository.TenderRepository
	employeeRepo repository.EmployeeRepository
	trManager    *manager.Manager
}

func NewBidService(repo repository.BidRepository,
	tenderRepo repository.TenderRepository,
	employeeRepo repository.EmployeeRepository,
	trManager *manager.Manager) *bidService {
	return &bidService{
		bidRepo:      repo,
		employeeRepo: employeeRepo,
		tenderRepo:   tenderRepo,
		trManager:    trManager,
	}
}

func (s *bidService) GetMy(ctx context.Context, limit, offset uint, username string) ([]model.Bid, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return s.bidRepo.GetMy(ctx, limit, offset, employee.ID)
}

func (s *bidService) GetTenderList(ctx context.Context, tenderID uuid.UUID,
	limit, offset uint, username string) ([]model.Bid, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}

	_, err = s.tenderRepo.GetById(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	return s.bidRepo.GetTenderBidList(ctx, tenderID, limit, offset, employee.ID, organization.ID)
}

func (s *bidService) GetStatus(ctx context.Context, bidID uuid.UUID, username string) (enum.BidStatus, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	statusResponse, err := s.bidRepo.GetStatus(ctx, bidID)
	if err != nil {
		return "", err
	}

	organization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return "", errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return "", err
	}

	authorOrganization, err := s.employeeRepo.GetUserOrganizationById(ctx, statusResponse.AuthorID)
	if err != nil {
		return "", err
	}

	isAuthor := statusResponse.AuthorID == employee.ID
	isFromAuthorOrganization := authorOrganization.ID == organization.ID
	isFromTenderOrganization := statusResponse.OrganizationID == organization.ID
	isTenderOrganizationHasAccess := statusResponse.Status == enum.BidPublished && isFromTenderOrganization

	if !isAuthor && !isFromAuthorOrganization && !isTenderOrganizationHasAccess {
		return "", errors.Wrap(model.ErrForbidden, "access denied to bid status")
	}

	return statusResponse.Status, nil
}

func (s *bidService) GetTenderReviews(ctx context.Context, limit, offset uint,
	tenderID uuid.UUID, author, requester string) ([]model.BidReview, error) {
	if _, err := s.tenderRepo.GetById(ctx, tenderID); err != nil {
		return nil, err
	}

	employeeAuthor, err := s.employeeRepo.GetByUsername(ctx, author)
	if err != nil {
		return nil, err
	}

	if _, err = s.employeeRepo.GetByUsername(ctx, requester); err != nil {
		return nil, err
	}

	authorOrganization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, author)
	if err != nil {
		return nil, err
	}

	requesterOrganization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, requester)
	if err != nil {
		return nil, err
	}
	if authorOrganization.ID != requesterOrganization.ID {
		return nil, errors.Wrap(model.ErrForbidden, "access denied")
	}

	return s.bidRepo.GetTenderReviews(ctx, limit, offset, tenderID, employeeAuthor.ID, requesterOrganization.ID)
}

func (s *bidService) Create(ctx context.Context, entity model.Bid) (*model.Bid, error) {
	employee, err := s.employeeRepo.GetById(ctx, entity.AuthorID)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}
	if organization == nil {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	_, err = s.tenderRepo.GetById(ctx, entity.TenderID)
	if err != nil {
		return nil, err
	}

	return s.bidRepo.Create(ctx, entity)
}

func (s *bidService) Edit(ctx context.Context, entity model.Bid, username string) (*model.Bid, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	bid, err := s.bidRepo.GetById(ctx, entity.ID)
	if err != nil {
		return nil, err
	}
	if bid.AuthorID != employee.ID {
		return nil, errors.Wrap(model.ErrForbidden, "access denied")
	}

	return s.bidRepo.Edit(ctx, entity)
}

func (s *bidService) UpdateStatus(ctx context.Context, bidID uuid.UUID, status enum.BidStatus, username string) (*model.Bid, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}

	bid, err := s.bidRepo.GetById(ctx, bidID)
	if err != nil {
		return nil, err
	}

	ownerId, err := s.bidRepo.GetTenderOrganizationId(ctx, bidID)
	if err != nil {
		return nil, err
	}

	if employee.ID != bid.AuthorID && organization.ID != ownerId {
		return nil, errors.Wrap(model.ErrForbidden, "access denied")
	}

	return s.bidRepo.UpdateStatus(ctx, bidID, status)
}

func (s *bidService) SubmitDecision(ctx context.Context, bidID uuid.UUID, decision enum.BidDecision, username string) (*model.Bid, error) {
	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}

	bid, err := s.bidRepo.GetById(ctx, bidID)
	if err != nil {
		return nil, err
	}

	ownerId, err := s.bidRepo.GetTenderOrganizationId(ctx, bidID)
	if err != nil {
		return nil, err
	}

	if employee.ID != bid.AuthorID && organization.ID != ownerId {
		return nil, errors.Wrap(model.ErrForbidden, "access denied")
	}

	err = s.trManager.Do(ctx, func(ctx context.Context) error {
		if err := s.bidRepo.SubmitDecision(ctx, bidID, employee.ID, decision); err != nil {
			return errors.Wrap(err, "submit decision")
		}

		if decision == enum.BidRejected {
			_, err = s.bidRepo.UpdateStatus(ctx, bidID, enum.BidCanceled)
			if err != nil {
				return errors.Wrap(err, "cancel bid")
			}
			return nil
		}

		decisions, err := s.bidRepo.GetDecisions(ctx, bidID)
		if err != nil {
			return errors.Wrap(err, "get decisions")
		}

		quorum, err := s.bidRepo.GetQuorum(ctx, bidID)
		if err != nil {
			return errors.Wrap(err, "get quorum")
		}

		if len(decisions) >= quorum {
			_, err = s.bidRepo.UpdateStatus(ctx, bidID, enum.BidCanceled)
			if err != nil {
				return errors.Wrap(err, "approve bid")
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "transaction failed")
	}

	return s.bidRepo.GetById(ctx, bidID)
}

func (s *bidService) Feedback(ctx context.Context, bidID uuid.UUID, feedback, username string) (*model.Bid, error) {
	ownerId, err := s.bidRepo.GetTenderOrganizationId(ctx, bidID)
	if err != nil {
		return nil, err
	}

	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	organization, err := s.employeeRepo.GetUserOrganizationByUsername(ctx, employee.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(model.ErrForbidden, "user has no access")
		}
		return nil, err
	}
	if organization.ID != ownerId {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	if err := s.bidRepo.Feedback(ctx, bidID, feedback); err != nil {
		return nil, err
	}

	return s.bidRepo.GetById(ctx, bidID)
}

func (s *bidService) Rollback(ctx context.Context, bidID uuid.UUID, version uint64, username string) (*model.Bid, error) {
	bidToRollback, err := s.bidRepo.GetByVersion(ctx, bidID, version)
	if err != nil {
		return nil, err
	}

	employee, err := s.employeeRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if employee.ID != bidToRollback.AuthorID {
		return nil, errors.Wrap(model.ErrForbidden, "user has no access")
	}

	updated, err := s.bidRepo.Edit(ctx, *bidToRollback)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
