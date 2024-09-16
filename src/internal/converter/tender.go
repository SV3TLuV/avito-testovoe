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

func ToTenderHistoryRecordFromTenderHistory(tender model.TenderHistory) goqu.Record {
	record := goqu.Record{}

	if tender.ID != uuid.Nil {
		record["id"] = tender.ID
	}
	if tender.TenderID != uuid.Nil {
		record["tender_id"] = tender.TenderID
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

func ToTenderHistoryFromTender(tender model.Tender) model.TenderHistory {
	return model.TenderHistory{
		TenderID:       tender.ID,
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType,
		Status:         tender.Status,
		OrganizationID: tender.OrganizationID,
		Version:        tender.Version,
		CreatedAt:      tender.CreatedAt,
		UpdatedAt:      tender.UpdatedAt,
	}
}

func ToTenderFromTenderHistory(history model.TenderHistory) model.Tender {
	return model.Tender{
		ID:             history.TenderID,
		Name:           history.Name,
		Description:    history.Description,
		ServiceType:    history.ServiceType,
		Status:         history.Status,
		OrganizationID: history.OrganizationID,
		Version:        history.Version,
		CreatedAt:      history.CreatedAt,
		UpdatedAt:      history.UpdatedAt,
	}
}

func ToTenderViewFromTender(tender model.Tender) model.TenderView {
	return model.TenderView{
		ID:          tender.ID,
		Name:        tender.Name,
		Description: tender.Description,
		Status:      tender.Status,
		ServiceType: tender.ServiceType,
		Version:     tender.Version,
		CreatedAt:   tender.CreatedAt,
	}
}

func ToTenderViewsFromTender(tenders []model.Tender) []model.TenderView {
	views := make([]model.TenderView, 0, len(tenders))
	for _, tender := range tenders {
		views = append(views, ToTenderViewFromTender(tender))
	}
	return views
}
