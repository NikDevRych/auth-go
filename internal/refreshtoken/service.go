package refreshtoken

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) NewRefreshToken(ctx context.Context, userId string) (*RefreshToken, error) {
	token := New(userId)
	if err := s.repo.CreateOrUpdate(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}
