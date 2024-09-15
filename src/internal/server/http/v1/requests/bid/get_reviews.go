package bid

import "github.com/google/uuid"

type GetReviewsRequest struct {
	TenderID          uuid.UUID `json:"tenderId" validate:"required,uuid,max=100"`
	AuthorUsername    string    `query:"authorUsername" validate:"required,max=50"`
	RequesterUsername string    `query:"requesterUsername" validate:"required,max=50"`
	Limit             int64     `query:"limit" validate:"gte=0"`
	Offset            int64     `query:"offset" validate:"gte=0"`
}

func (r *GetReviewsRequest) SetDefaults() {
	if r.Limit == 0 {
		r.Limit = 5
	}
	if r.Offset == 0 {
		r.Offset = 0
	}
}
