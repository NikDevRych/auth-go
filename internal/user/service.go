package user

import (
	"context"
	"errors"

	"github.com/NikDevRych/auth-go/internal/auth"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) SignUp(ctx context.Context, req *UserDataRequest) error {
	user, err := New(req.Email, req.Password)
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, user)
}

func (s *service) SignIn(ctx context.Context, req *UserDataRequest) (string, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	if !user.isPasswordMatch(req.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := auth.CreateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
