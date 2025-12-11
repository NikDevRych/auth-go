package db

import (
	"context"

	"github.com/NikDevRych/auth-go/internal/refreshtoken"
	"github.com/jackc/pgx/v5/pgxpool"
)

type refreshTokenRepository struct {
	dbpool *pgxpool.Pool
}

func NewRefreshTokenRepository(dbpool *pgxpool.Pool) *refreshTokenRepository {
	return &refreshTokenRepository{dbpool: dbpool}
}

func (r *refreshTokenRepository) CreateOrUpdate(ctx context.Context, token *refreshtoken.RefreshToken) error {
	const query = `
		INSERT INTO refresh_tokens (user_id, token, expire_at, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id)
		DO UPDATE SET
			token = EXCLUDED.token,
			expire_at = EXCLUDED.expire_at,
			created_at = EXCLUDED.created_at;
	`

	_, err := r.dbpool.Exec(ctx, query, token.UserId, token.Token, token.ExpireAt, token.CreateAt)
	if err != nil {
		return err
	}

	return nil
}
