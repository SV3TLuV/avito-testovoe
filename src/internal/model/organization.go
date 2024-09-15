package model

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
	"time"
)

type Organization struct {
	ID          uuid.UUID
	Name        string
	Description string
	Type        enum.OrganizationType
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
