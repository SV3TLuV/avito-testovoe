package enum

type AuthorType string

const (
	Organization AuthorType = "Organization"
	User         AuthorType = "User"
)

func (val AuthorType) IsValid() bool {
	switch val {
	case Organization, User:
		return true
	default:
		return false
	}
}
