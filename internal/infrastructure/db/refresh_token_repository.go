package db

import (
	"context"

	"github.com/NikDevRych/auth-go/internal/refreshtoken"
	"github.com/jackc/pgx/v5"
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

func (r *refreshTokenRepository) FindByToken(ctx context.Context, token string) (*refreshtoken.RefreshToken, error) {
	const query = `
		SELECT user_id, token, expire_at, created_at
		FROM refresh_tokens
		WHERE token = $1
	`

	var refreshToken refreshtoken.RefreshToken
	err := r.dbpool.QueryRow(ctx, query, token).Scan(&refreshToken.UserId, &refreshToken.Token, &refreshToken.ExpireAt, &refreshToken.CreateAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &refreshToken, nil
}
