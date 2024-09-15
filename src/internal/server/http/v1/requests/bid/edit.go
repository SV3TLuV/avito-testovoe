package bid

import (
	"github.com/google/uuid"
)

type EditRequest struct {
	BidID       uuid.UUID `json:"bidId" validate:"required,uuid,max=100"`
	Username    string    `query:"username" validate:"required,max=50"`
	Name        string    `body:"name" validate:"required,max=100"`
	Description string    `body:"description" validate:"required,max=500"`
}
