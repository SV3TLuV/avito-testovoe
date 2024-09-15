package bid

import "github.com/google/uuid"

type FeedbackRequest struct {
	BidID       uuid.UUID `json:"bidId" validate:"required,uuid,max=100"`
	Username    string    `query:"username" validate:"required,max=50"`
	BidFeedback string    `query:"bidFeedback" validate:"required,max=1000"`
}
