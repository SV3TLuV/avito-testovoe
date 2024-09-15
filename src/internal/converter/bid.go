package converter

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"tender_api/src/internal/model"
)

func ToBidRecordFromBid(bid model.Bid) goqu.Record {
	record := goqu.Record{}

	if bid.ID != uuid.Nil {
		record["id"] = bid.ID
	}
	if bid.Name != "" {
		record["name"] = bid.Name
	}
	if bid.Description != "" {
		record["description"] = bid.Description
	}
	if bid.Status != "" {
		record["status"] = bid.Status
	}
	if bid.Decision != "" {
		record["decision"] = bid.Decision
	}
	if bid.TenderID != uuid.Nil {
		record["tender_id"] = bid.TenderID
	}
	if bid.AuthorType != "" {
		record["author_type"] = bid.AuthorType
	}
	if bid.AuthorID != uuid.Nil {
		record["author_id"] = bid.AuthorID
	}
	if bid.Version != 0 {
		record["version"] = bid.Version
	}
	if !bid.CreatedAt.IsZero() {
		record["created_at"] = bid.CreatedAt
	}
	if !bid.UpdatedAt.IsZero() {
		record["updated_at"] = bid.UpdatedAt
	}

	return record
}
