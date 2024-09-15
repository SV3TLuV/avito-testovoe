package tender

import (
	"tender_api/src/internal/model/enum"
)

type GetListRequest struct {
	Limit       int64                    `query:"limit" validate:"gte=0"`
	Offset      int64                    `query:"offset" validate:"gte=0"`
	ServiceType []enum.TenderServiceType `query:"service_type" validate:"enum_tender_service_type"`
}

func (r *GetListRequest) SetDefaults() {
	if r.Limit == 0 {
		r.Limit = 5
	}
	if r.Offset == 0 {
		r.Offset = 0
	}
}
