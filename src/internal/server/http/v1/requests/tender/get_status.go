package tender

import "github.com/google/uuid"

type GetStatusRequest struct {
	TenderID uuid.UUID `json:"tenderId" validate:"required,uuid,max=100"`
	Username string    `query:"username" validate:"required,max=50"`
}
