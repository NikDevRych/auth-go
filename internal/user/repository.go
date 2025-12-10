package user

import "context"

type Repository interface {
	Create(context.Context, *User) error
	FindByEmail(context.Context, string) (*User, error)
}
