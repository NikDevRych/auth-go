package refreshtoken

import "context"

type Repository interface {
	CreateOrUpdate(context.Context, *RefreshToken) error
	FindByToken(context.Context, string) (*RefreshToken, error)
}
