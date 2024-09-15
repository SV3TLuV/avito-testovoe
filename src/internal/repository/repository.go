package repository

import (
	"context"
	"github.com/google/uuid"
	"tender_api/src/internal/model"
	"tender_api/src/internal/model/enum"
)

type TenderRepository interface {
	GetList(ctx context.Context, limit, offset uint, serviceType []enum.TenderServiceType) ([]model.Tender, error)
	GetMy(ctx context.Context, limit, offset uint, employeeID uuid.UUID) ([]model.Tender, error)
	GetStatus(ctx context.Context, tenderID uuid.UUID) (*model.TenderStatusResponse, error)
	GetById(ctx context.Context, tenderID uuid.UUID) (*model.Tender, error)
	GetByVersion(ctx context.Context, tenderID uuid.UUID, version uint64) (*model.Tender, error)
	Create(ctx context.Context, entity model.Tender, employeeID uuid.UUID) (*model.Tender, error)
	Edit(ctx context.Context, entity model.Tender) (*model.Tender, error)
	UpdateStatus(ctx context.Context, tenderID uuid.UUID, status enum.TenderStatus) (*model.Tender, error)
}

type BidRepository interface {
	GetMy(ctx context.Context, limit, offset uint, employeeId uuid.UUID) ([]model.Bid, error)
	GetTenderList(ctx context.Context, tenderID uuid.UUID, limit, offset uint,
		employeeID, organizationID uuid.UUID) ([]model.Bid, error)
	GetStatus(ctx context.Context, bidID uuid.UUID) (*model.BidStatusResponse, error)
	GetTenderReviews(ctx context.Context, limit, offset uint,
		tenderID, authorID, requesterOrganizationID uuid.UUID) ([]model.BidReview, error)
	GetById(ctx context.Context, bidID uuid.UUID) (*model.Bid, error)
	GetTenderOwnerId(ctx context.Context, bidID uuid.UUID) (uuid.UUID, error)
	GetByVersion(ctx context.Context, bidID uuid.UUID, version uint64) (*model.Bid, error)
	Create(ctx context.Context, entity model.Bid) (*model.Bid, error)
	Edit(ctx context.Context, entity model.Bid) (*model.Bid, error)
	UpdateStatus(ctx context.Context, bidID uuid.UUID, status enum.BidStatus) (*model.Bid, error)
	SubmitDecision(ctx context.Context, bidID uuid.UUID, decision enum.BidDecision) error
	Feedback(ctx context.Context, bidID uuid.UUID, feedback string) error
}

type EmployeeRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.Employee, error)
	GetUserOrganization(ctx context.Context, username string) (*model.Organization, error)
}
