package model

import "github.com/google/uuid"

type OrganizationResponsible struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	UserID         uuid.UUID
}
