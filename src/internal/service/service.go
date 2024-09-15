package service

import (
	"context"
	"github.com/google/uuid"
	"tender_api/src/internal/model"
	"tender_api/src/internal/model/enum"
)

type TenderService interface {
	GetList(ctx context.Context, limit, offset uint, serviceType []enum.TenderServiceType) ([]model.Tender, error)
	GetMy(ctx context.Context, limit, offset uint, username string) ([]model.Tender, error)
	GetStatus(ctx context.Context, tenderID uuid.UUID, username string) (enum.TenderStatus, error)
	Create(ctx context.Context, entity model.Tender, username string) (*model.Tender, error)
	Edit(ctx context.Context, entity model.Tender, username string) (*model.Tender, error)
	UpdateStatus(ctx context.Context, tenderID uuid.UUID,
		status enum.TenderStatus, username string) (*model.Tender, error)
	Rollback(ctx context.Context, tenderID uuid.UUID,
		version uint64, username string) (*model.Tender, error)
}

type BidService interface {
	GetMy(ctx context.Context, limit, offset uint, username string) ([]model.Bid, error)
	GetTenderList(ctx context.Context, tenderID uuid.UUID, limit, offset uint, username string) ([]model.Bid, error)
	GetStatus(ctx context.Context, bidID uuid.UUID, username string) (enum.BidStatus, error)
	GetTenderReviews(ctx context.Context, limit, offset uint,
		tenderID uuid.UUID, author, requester string) ([]model.BidReview, error)
	Create(ctx context.Context, entity model.Bid) (*model.Bid, error)
	Edit(ctx context.Context, entity model.Bid, username string) (*model.Bid, error)
	UpdateStatus(ctx context.Context, bidID uuid.UUID, status enum.BidStatus, username string) (*model.Bid, error)
	SubmitDecision(ctx context.Context, bidID uuid.UUID, decision enum.BidDecision, username string) (*model.Bid, error)
	Feedback(ctx context.Context, bidID uuid.UUID, feedback, username string) (*model.Bid, error)
	Rollback(ctx context.Context, bidID uuid.UUID, version uint64, username string) (*model.Bid, error)
}
