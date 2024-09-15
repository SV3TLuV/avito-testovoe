package enum

type BidStatus string

const (
	BidCreated   BidStatus = "Created"
	BidPublished BidStatus = "Published"
	BidCanceled  BidStatus = "Canceled"
)

func (val BidStatus) IsValid() bool {
	switch val {
	case BidCreated, BidPublished, BidCanceled:
		return true
	default:
		return false
	}
}
