package refreshtoken

import "context"

type Repository interface {
	CreateOrUpdate(context.Context, *RefreshToken) error
}
