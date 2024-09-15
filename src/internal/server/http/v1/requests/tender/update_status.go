package tender

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type UpdateStatusRequest struct {
	TenderID uuid.UUID         `json:"tenderId" validate:"required,uuid,max=100"`
	Status   enum.TenderStatus `query:"status" validate:"required,enum_tender_status"`
	Username string            `query:"username" validate:"required,max=50"`
}
