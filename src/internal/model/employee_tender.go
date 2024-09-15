package model

import "github.com/google/uuid"

type EmployeeTender struct {
	ID         uuid.UUID
	EmployeeID uuid.UUID
	TenderID   uuid.UUID
}
