package usecase

import (
	"context"
	"errors"

	"rubix-store/service/auth/domain"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserExists      = errors.New("user already exists")
	ErrHashPassword    = errors.New("failed to hash password")
)

type RegisterUsecase struct {
	repo domain.UserRepository
}

func NewRegisterUsecase(repo domain.UserRepository) *RegisterUsecase {
	return &RegisterUsecase{repo: repo}
}

func (uc *RegisterUsecase) Execute(ctx context.Context, email, password, roles string) (*domain.User, error) {
	// Check if user already exists
	existing, _ := uc.repo.GetByEmail(ctx, email)
	if existing != nil {
		return nil, ErrUserExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrHashPassword
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: string(hash),
		Roles:        roles,
	}

	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
