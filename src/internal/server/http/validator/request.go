package validator

import (
	"github.com/go-playground/validator"
)

type requestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *requestValidator {
	v := validator.New()
	_ = v.RegisterValidation("uuid", validateUUID)
	_ = v.RegisterValidation("enum_author_type", validateAuthorType)
	_ = v.RegisterValidation("enum_bid_decision", validateBidDecision)
	_ = v.RegisterValidation("enum_bid_status", validateBidStatus)
	_ = v.RegisterValidation("enum_organization_type", validateOrganizationType)
	_ = v.RegisterValidation("enum_tender_service_type", validateTenderServiceType)
	_ = v.RegisterValidation("enum_tender_status", validateTenderStatus)
	return &requestValidator{
		validator: v,
	}
}

func (v *requestValidator) Validate(i any) error {
	return v.validator.Struct(i)
}
