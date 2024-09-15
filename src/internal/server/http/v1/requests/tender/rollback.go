package tender

import "github.com/google/uuid"

type RollbackRequest struct {
	TenderID uuid.UUID `json:"tenderId" validate:"required,uuid,max=100"`
	Version  uint64    `json:"version" validate:"required,gte=1"`
	Username string    `query:"username" validate:"required,max=50"`
}
