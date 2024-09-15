package converter

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"tender_api/src/internal/model"
)

func ToTenderRecordFromTender(tender model.Tender) goqu.Record {
	record := goqu.Record{}

	if tender.ID != uuid.Nil {
		record["id"] = tender.ID
	}
	if tender.Name != "" {
		record["name"] = tender.Name
	}
	if tender.Description != "" {
		record["description"] = tender.Description
	}
	if tender.ServiceType != "" {
		record["service_type"] = tender.ServiceType
	}
	if tender.Status != "" {
		record["status"] = tender.Status
	}
	if tender.OrganizationID != uuid.Nil {
		record["organization_id"] = tender.OrganizationID
	}
	if tender.Version != 0 {
		record["version"] = tender.Version
	}
	if !tender.CreatedAt.IsZero() {
		record["created_at"] = tender.CreatedAt
	}
	if !tender.UpdatedAt.IsZero() {
		record["updated_at"] = tender.UpdatedAt
	}

	return record
}
