package bid

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
)

type CreateRequest struct {
	Name        string          `body:"name" validate:"required,max=100"`
	Description string          `body:"description" validate:"required,max=500"`
	TenderID    uuid.UUID       `body:"tenderId" validate:"required,uuid,max=100"`
	AuthorType  enum.AuthorType `body:"authorType" validate:"required,enum_author_type"`
	AuthorID    uuid.UUID       `body:"authorId" validate:"required,uuid,max=100"`
}
