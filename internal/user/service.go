package user

import "context"

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
