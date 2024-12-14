package login

import "context"

type Service interface {
	ValidateCredentials(ctx context.Context, username, password string) (bool, error)
}
