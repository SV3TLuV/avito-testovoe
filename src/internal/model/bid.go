package model

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
	"time"
)

type Bid struct {
	ID          uuid.UUID
	Name        string
	Description string
	Status      enum.BidStatus
	Decision    enum.BidDecision
	TenderID    uuid.UUID
	AuthorType  enum.AuthorType
	AuthorID    uuid.UUID
	Version     uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BidView struct {
	ID         uuid.UUID       `json:"id"`
	Name       string          `json:"name"`
	Status     enum.BidStatus  `json:"status"`
	AuthorType enum.AuthorType `json:"authorType"`
	AuthorID   uuid.UUID       `json:"authorId"`
	Version    uint64          `json:"version"`
	CreatedAt  time.Time       `json:"createdAt"`
}

type BidStatusResponse struct {
	Status         enum.BidStatus
	OrganizationID uuid.UUID
	AuthorID       uuid.UUID
}
