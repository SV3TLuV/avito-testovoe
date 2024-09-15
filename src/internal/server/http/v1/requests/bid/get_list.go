package bid

import "github.com/google/uuid"

type GetListRequest struct {
	TenderID uuid.UUID `json:"tenderId" validate:"required,uuid,max=100"`
	Username string    `query:"username" validate:"required,max=50"`
	Limit    int64     `query:"limit" validate:"gte=0"`
	Offset   int64     `query:"offset" validate:"gte=0"`
}

func (r *GetListRequest) SetDefaults() {
	if r.Limit == 0 {
		r.Limit = 5
	}
	if r.Offset == 0 {
		r.Offset = 0
	}
}
