package domain

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
