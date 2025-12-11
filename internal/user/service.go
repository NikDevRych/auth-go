package user

import (
	"context"
	"errors"
	"time"

	"github.com/NikDevRych/auth-go/internal/auth"
	"github.com/NikDevRych/auth-go/internal/config"
	"github.com/NikDevRych/auth-go/internal/refreshtoken"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type service struct {
	cfg        *config.Config
	repo       Repository
	refreshSvc *refreshtoken.Service
}

func NewService(cfg *config.Config, userRepo Repository, refreshSvc *refreshtoken.Service) *service {
	return &service{
		cfg:        cfg,
		repo:       userRepo,
		refreshSvc: refreshSvc,
	}
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

	accessToken, err := auth.CreateToken(s.cfg.JWTSecretKey)
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

func (s *service) RefreshAccessToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.TokenResponse, error) {
	token, err := s.refreshSvc.FindByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if token == nil || token.ExpireAt.Before(time.Now().UTC()) {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := auth.CreateToken(s.cfg.JWTSecretKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.refreshSvc.NewRefreshToken(ctx, token.UserId)
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
