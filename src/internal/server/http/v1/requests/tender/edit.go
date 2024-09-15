package tender

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type EditRequest struct {
	TenderID    uuid.UUID              `json:"tenderId" validate:"required,uuid,max=100"`
	Username    string                 `query:"username" validate:"required,max=50"`
	Name        string                 `body:"name" validate:"required,max=100"`
	Description string                 `body:"description" validate:"required,max=500"`
	ServiceType enum.TenderServiceType `body:"serviceType" validate:"required,enum_tender_service_type"`
}
