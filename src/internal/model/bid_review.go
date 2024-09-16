package model

import (
	"github.com/google/uuid"
	"time"
)

type BidReview struct {
	ID          uuid.UUID
	BidID       uuid.UUID
	EmployeeID  uuid.UUID
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BidReviewView struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
