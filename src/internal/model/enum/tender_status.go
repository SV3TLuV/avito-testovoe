package enum

type TenderStatus string

const (
	TenderCreated   TenderStatus = "Created"
	TenderPublished TenderStatus = "Published"
	TenderClosed    TenderStatus = "Closed"
)

func (val TenderStatus) IsValid() bool {
	switch val {
	case TenderCreated, TenderPublished, TenderClosed:
		return true
	default:
		return false
	}
}
