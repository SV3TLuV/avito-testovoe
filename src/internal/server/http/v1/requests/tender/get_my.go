package tender

type GetMyRequest struct {
	Limit    int64  `query:"limit" validate:"gte=0"`
	Offset   int64  `query:"offset" validate:"gte=0"`
	Username string `query:"username" validate:"required,max=50"`
}

func (r *GetMyRequest) SetDefaults() {
	if r.Limit == 0 {
		r.Limit = 5
	}
	if r.Offset == 0 {
		r.Offset = 0
	}
}
