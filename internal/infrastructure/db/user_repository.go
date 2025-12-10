package db

import (
	"context"

	"github.com/NikDevRych/auth-go/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	dbpool *pgxpool.Pool
}

func NewUserRepository(dbpool *pgxpool.Pool) *userRepository {
	return &userRepository{dbpool: dbpool}
}

func (r *userRepository) Create(ctx context.Context, user *user.User) error {
	const query = `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
	`

	_, err := r.dbpool.Exec(ctx, query, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	const query = `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var user user.User
	err := r.dbpool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
