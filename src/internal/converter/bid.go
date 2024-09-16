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

func ToBidHistoryRecordFromBidHistory(bid model.BidHistory) goqu.Record {
	record := goqu.Record{}

	if bid.ID != uuid.Nil {
		record["id"] = bid.ID
	}
	if bid.BidID != uuid.Nil {
		record["bid_id"] = bid.BidID
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

func ToBidHistoryFromBid(bid model.Bid) model.BidHistory {
	return model.BidHistory{
		BidID:       bid.ID,
		Name:        bid.Name,
		Description: bid.Description,
		Status:      bid.Status,
		Decision:    bid.Decision,
		TenderID:    bid.TenderID,
		AuthorType:  bid.AuthorType,
		AuthorID:    bid.AuthorID,
		Version:     bid.Version,
		CreatedAt:   bid.CreatedAt,
		UpdatedAt:   bid.UpdatedAt,
	}
}

func ToBidFromBidHistory(history model.BidHistory) model.Bid {
	return model.Bid{
		ID:          history.BidID,
		Name:        history.Name,
		Description: history.Description,
		Status:      history.Status,
		Decision:    history.Decision,
		TenderID:    history.TenderID,
		AuthorType:  history.AuthorType,
		AuthorID:    history.AuthorID,
		Version:     history.Version,
		CreatedAt:   history.CreatedAt,
		UpdatedAt:   history.UpdatedAt,
	}
}

func ToBidViewFromBid(bid model.Bid) model.BidView {
	return model.BidView{
		ID:         bid.ID,
		Name:       bid.Name,
		Status:     bid.Status,
		AuthorType: bid.AuthorType,
		AuthorID:   bid.AuthorID,
		Version:    bid.Version,
		CreatedAt:  bid.CreatedAt,
	}
}

func ToBidViewsFromBid(bids []model.Bid) []model.BidView {
	views := make([]model.BidView, 0, len(bids))
	for _, bid := range bids {
		views = append(views, ToBidViewFromBid(bid))
	}
	return views
}

func ToBidReviewViewFromBidReview(review model.BidReview) model.BidReviewView {
	return model.BidReviewView{
		ID:          review.ID,
		Description: review.Description,
		CreatedAt:   review.CreatedAt,
	}
}

func ToBidReviewViewsFromBidReview(reviews []model.BidReview) []model.BidReviewView {
	views := make([]model.BidReviewView, 0, len(reviews))
	for _, review := range reviews {
		views = append(views, ToBidReviewViewFromBidReview(review))
	}
	return views
}
