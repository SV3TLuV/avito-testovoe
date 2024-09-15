package bid

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type UpdateStatusRequest struct {
	BidID    uuid.UUID      `json:"bidId" validate:"required,uuid,max=100"`
	Status   enum.BidStatus `query:"status" validate:"required,enum_bid_status"`
	Username string         `query:"username" validate:"required,max=50"`
}
