package game

// ContextKey constants.
const (
	CorrelationIdContext ContextKey = "correlation-id"
	ServiceNameContext   ContextKey = "service-name"
)

// ContextKey is a custom type for context keys.
type ContextKey string

func (k ContextKey) String() string {
	return string(k)
}
