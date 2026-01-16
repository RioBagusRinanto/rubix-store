package service

import (
	"errors"
	"rubix-store/internal/models"
	"rubix-store/internal/repository"
	"rubix-store/internal/utils"
)

type AuthService interface {
	Register(email, password, firstName, lastName string) (*models.User, error)
	Login(email, password string) (string, *models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(email, password, firstName, lastName string) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	user := &models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      "customer",
		IsActive:  true,
	}

	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(email, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", nil, errors.New("account is deactivated")
	}

	if !user.CheckPassword(password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *authService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *authService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}
