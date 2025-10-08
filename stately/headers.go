package stately

const (
	RequestIDHeader = "X-Stately-Request-Id"
	NoAdminHeader   = "X-Stately-NoAdmin"
	StoreIDHeader   = "X-Stately-StoreId"
)

var AllCustomHeaders = []string{
	RequestIDHeader,
	NoAdminHeader,
	StoreIDHeader,
}
