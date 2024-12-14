package logger

// ErrorField creates a new field with the given error.
func ErrorField(value error) LogField { return NewNonSensitiveField("error", value) }
