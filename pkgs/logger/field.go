package logger

import "go.uber.org/zap"

type Field interface {
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
func NewField[T any](key string, val T) Field {
	return field[T]{key: key, value: DefaultMasker()(val)}
}

// newNonSensitiveField returns a new field with the given key and value and not masked.
func newNonSensitiveField[T any](key string, val T) Field {
	return field[T]{key: key, value: val}
}

// newSensitiveField returns a new field with the given key with value masked by the given masker.
func newSensitiveField[T any](key string, val T, mask Masker[T]) Field {
	if mask == nil {
		return NewField[T](key, val)
	}

	return field[T]{key: key, value: mask(val)}
}
