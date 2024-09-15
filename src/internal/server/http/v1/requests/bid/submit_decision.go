package bid

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type SubmitDecisionRequest struct {
	BidID    uuid.UUID        `json:"bidId" validate:"required,uuid,max=100"`
	Decision enum.BidDecision `query:"decision" validate:"required,enum_bid_decision"`
	Username string           `query:"username" validate:"required,max=50"`
}
