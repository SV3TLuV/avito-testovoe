package bid

import "github.com/google/uuid"

type GetStatusRequest struct {
	BidID    uuid.UUID `json:"bidId" validate:"required,uuid,max=100"`
	Username string    `query:"username" validate:"required,max=50"`
}
