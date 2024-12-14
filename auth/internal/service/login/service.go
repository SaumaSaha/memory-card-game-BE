package login

import "context"

// Service ...
type Service interface {
	ValidateCredentials(ctx context.Context, username, password string) (bool, error)
}
