package bid

type GetMyRequest struct {
	Username string `query:"username" validate:"required,max=50"`
	Limit    uint   `query:"limit" validate:"gte=0"`
	Offset   uint   `query:"offset" validate:"gte=0"`
}

func (r *GetMyRequest) SetDefaults() {
	if r.Limit == 0 {
		r.Limit = 5
	}
	if r.Offset == 0 {
		r.Offset = 0
	}
}
