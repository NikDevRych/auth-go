package user

import (
	"context"
	"errors"

	"github.com/NikDevRych/auth-go/internal/auth"
	"github.com/NikDevRych/auth-go/internal/refreshtoken"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type service struct {
	repo       Repository
	refreshSvc *refreshtoken.Service
}

func NewService(userRepo Repository, refreshSvc *refreshtoken.Service) *service {
	return &service{repo: userRepo, refreshSvc: refreshSvc}
}

func (s *service) SignUp(ctx context.Context, req *UserDataRequest) error {
	user, err := New(req.Email, req.Password)
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, user)
}

func (s *service) SignIn(ctx context.Context, req *UserDataRequest) (*auth.TokenResponse, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !user.isPasswordMatch(req.Password) {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := auth.CreateToken(user.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.refreshSvc.NewRefreshToken(ctx, user.ID.String())
	if err != nil {
		return nil, err
	}

	tokenResponse := &auth.TokenResponse{
		AccessToken: accessToken,
		RefreshToken: auth.RefreshTokenResponse{
			RefreshToken: refreshToken.Token,
			ExpireAt:     refreshToken.ExpireAt,
		},
	}

	return tokenResponse, nil
}
