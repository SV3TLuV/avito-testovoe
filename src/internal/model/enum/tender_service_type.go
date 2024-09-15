package enum

type TenderServiceType string

const (
	Construction TenderServiceType = "Construction"
	Delivery     TenderServiceType = "Delivery"
	Manufacture  TenderServiceType = "Manufacture"
)

func (val TenderServiceType) IsValid() bool {
	switch val {
	case Construction, Delivery, Manufacture:
		return true
	default:
		return false
	}
}
