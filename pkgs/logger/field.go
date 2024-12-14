package logger

import "go.uber.org/zap"

// LogField is an interface that defines the ZapField method.
type LogField interface {
	ZapField() zap.Field
}

// field is a struct that implements the zap.Field
// interface. It is used to convert a field into a
// zap.Field.
type field[T any] struct {
	key   string
	value any
}

// ZapField converts a field into a zap.Field.
func (f field[T]) ZapField() zap.Field {
	return zap.Any(f.key, f.value)
}

// NewField returns a new field with the given key and value and not masked.
func NewField[T any](key string, val T) LogField {
	return field[T]{key: key, value: DefaultMasker()(val)}
}

// NewNonSensitiveField returns a new field with the given key and value and not masked.
func NewNonSensitiveField[T any](key string, val T) LogField {
	return field[T]{key: key, value: val}
}

// NewSensitiveField returns a new field with the given key with value masked by the given masker.
func NewSensitiveField[T any](key string, val T, mask Masker[T]) LogField {
	if mask == nil {
		return NewField[T](key, val)
	}

	return field[T]{key: key, value: mask(val)}
}
