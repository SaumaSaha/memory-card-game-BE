package game

const (
	CorrelationIdContext ContextKey = "correlation-id"
	ServiceNameContext   ContextKey = "service-name"
)

type ContextKey string

func (k ContextKey) String() string {
	return string(k)
}
