package enum

type BidDecision string

const (
	BidApproved BidDecision = "Approved"
	BidRejected BidDecision = "Rejected"
)

func (val BidDecision) IsValid() bool {
	switch val {
	case BidApproved, BidRejected:
		return true
	default:
		return false
	}
}
