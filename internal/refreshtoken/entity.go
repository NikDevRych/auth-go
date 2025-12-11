package refreshtoken

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	UserId   string
	Token    string
	ExpireAt time.Time
	CreateAt time.Time
}

func New(userId string) *RefreshToken {
	return &RefreshToken{
		UserId:   userId,
		Token:    uuid.NewString(),
		ExpireAt: time.Now().UTC().Add(30 * 24 * time.Hour),
		CreateAt: time.Now().UTC(),
	}
}
