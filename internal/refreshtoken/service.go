package refreshtoken

import "context"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) NewRefreshToken(ctx context.Context, userId string) (*RefreshToken, error) {
	token := New(userId)
	if err := s.repo.CreateOrUpdate(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}
