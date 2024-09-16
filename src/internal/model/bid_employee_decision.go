package model

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type BidEmployeeDecision struct {
	ID         uuid.UUID
	BidID      uuid.UUID
	EmployeeID uuid.UUID
	Decision   enum.BidDecision
}
