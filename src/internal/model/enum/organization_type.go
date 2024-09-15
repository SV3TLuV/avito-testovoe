package enum

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

func (val OrganizationType) IsValid() bool {
	switch val {
	case IE, LLC, JSC:
		return true
	default:
		return false
	}
}
