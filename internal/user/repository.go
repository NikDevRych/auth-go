package user

import "context"

type Repository interface {
	Create(context.Context, *User) error
}
