package logger

import (
	"fmt"
	"reflect"

	"github.com/ggwhite/go-masker"
)

// Masker is a function type that takes a value of type T and returns an any type.
type Masker[T any] func(T) any

// DefaultMasker returns a Masker that always masks the value,
// it always returns the value as ********.
func DefaultMasker() Masker[any] {
	return func(_ any) any { return "*********" }
}

// NameMasker returns a Masker that masks the name,
// it masks the second letter and the third letter.
func NameMasker() Masker[string] {
	return func(value string) any {
		return masker.New().Name(value)
	}
}

// EmailMasker returns a Masker that masks the email,
// it keeps the domain and the first 3 letters.
func EmailMasker() Masker[string] {
	return func(value string) any {
		return masker.New().Email(value)
	}
}

// MobileMasker returns a Masker that masks the mobile,
// it masks 3 digits from the 4'th digit.
func MobileMasker() Masker[string] {
	return func(value string) any {
		return masker.New().Mobile(value)
	}
}

// AddressMasker returns a Masker that masks the address,
// it keeps first 6 letters, mask the rest.
func AddressMasker() Masker[string] {
	return func(value string) any {
		return masker.New().Address(value)
	}
}

// PasswordMasker returns a Masker that masks the password,
// it always return ************.
func PasswordMasker() Masker[string] {
	return func(value string) any {
		return masker.New().Password(value)
	}
}

// CreditCardMasker returns a Masker that masks the credit card,
// it masks 6 digits from the 7'th digit.
func CreditCardMasker() Masker[string] {
	return func(value string) any {
		return masker.New().CreditCard(value)
	}
}

// TelephoneMasker returns a Masker that masks the telephone,
// it remove (, ),  , - chart, and mask last 4 digits of telephone number,
// format to (??)????-????.
func TelephoneMasker() Masker[string] {
	return func(value string) any {
		return masker.New().Telephone(value)
	}
}

// URLMasker returns a Masker that masks the url,
// it masks the password field if present.
func URLMasker() Masker[string] {
	return func(value string) any {
		return masker.New().URL(value)
	}
}

// StructMasker returns a Masker that masks the struct,
// it masks the password field if present.
func StructMasker() Masker[any] {
	maskStruct := func(value any) any {
		result, err := masker.New().Struct(value)
		if err != nil {
			return DefaultMasker()(value)
		}

		return result
	}

	return func(value any) any {
		if value == nil {
			return "nil"
		}

		vType := reflect.TypeOf(value)

		switch vType.Kind() {
		case reflect.Struct:
			return maskStruct(value)
		case reflect.Ptr:
			if vType.Elem().Kind() == reflect.Struct {
				return maskStruct(value)
			}

			return DefaultMasker()(value)
		default:
			return DefaultMasker()(value)
		}
	}
}

// IDMasker returns a Masker that masks the id,
// it masks the last 4 digits of ID number.
func IDMasker[T any]() Masker[T] {
	return func(value T) any {
		return masker.New().ID(fmt.Sprintf("%v", value))
	}
}
