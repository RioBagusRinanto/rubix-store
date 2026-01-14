package usecase

import (
	"context"
	"time"

	"rubix-store/service/auth/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	repo      domain.UserRepository
	jwtSecret []byte
}

func NewLoginUsecase(repo domain.UserRepository, jwtSecret []byte) *LoginUsecase {
	return &LoginUsecase{repo: repo, jwtSecret: jwtSecret}
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (uc *LoginUsecase) Execute(ctx context.Context, email, password string) (*LoginResponse, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidPassword
	}

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"roles": user.Roles,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(uc.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: signed}, nil
}
