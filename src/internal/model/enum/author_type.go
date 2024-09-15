package enum

type AuthorType string

const (
	Organization AuthorType = "organization"
	User         AuthorType = "user"
)

func (val AuthorType) IsValid() bool {
	switch val {
	case Organization, User:
		return true
	default:
		return false
	}
}
