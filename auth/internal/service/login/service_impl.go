package login

import (
	"context"
	"fmt"
)

type loginService struct{}

func NewService() Service {
	return &loginService{}
}

func (ls *loginService) ValidateCredentials(ctx context.Context, username, password string) (bool, error) {
	fmt.Printf("validating credentials %s %s\n", username, password)
	return true, nil
}
