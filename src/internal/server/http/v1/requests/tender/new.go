package tender

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type CreateRequest struct {
	Name           string                 `body:"name" validate:"required,max=100"`
	Description    string                 `body:"description" validate:"required,max=500"`
	ServiceType    enum.TenderServiceType `body:"serviceType" validate:"required,enum_tender_service_type"`
	OrganizationID uuid.UUID              `body:"organizationId" validate:"required,uuid,max=100"`
	Username       string                 `body:"creatorUsername" validate:"required,max=50"`
}
