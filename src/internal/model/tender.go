package model

import (
	"github.com/google/uuid"
	"tender_api/src/internal/model/enum"
	"time"
)

type Tender struct {
	ID             uuid.UUID
	Name           string
	Description    string
	ServiceType    enum.TenderServiceType
	Status         enum.TenderStatus
	OrganizationID uuid.UUID
	Version        uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TenderHistory struct {
	ID             uuid.UUID
	TenderID       uuid.UUID
	Name           string
	Description    string
	ServiceType    enum.TenderServiceType
	Status         enum.TenderStatus
	OrganizationID uuid.UUID
	Version        uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TenderView struct {
	ID          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Status      enum.TenderStatus      `json:"status"`
	ServiceType enum.TenderServiceType `json:"serviceType"`
	Version     uint64                 `json:"version"`
	CreatedAt   time.Time              `json:"createdAt"`
}

type TenderStatusResponse struct {
	Status         enum.TenderStatus
	OrganizationID uuid.UUID
}
