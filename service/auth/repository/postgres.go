package repository

import (
	"context"

	"rubix-store/service/auth/domain"

	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) domain.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, u *domain.User) error {
	_, err := r.db.NamedExecContext(ctx, `INSERT INTO users (email, password_hash, roles) VALUES (:email, :password_hash, :roles)`, u)
	return err
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	err := r.db.GetContext(ctx, &u, "SELECT id, email, password_hash, roles FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
