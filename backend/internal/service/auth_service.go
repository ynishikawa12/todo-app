package service

import (
	"errors"
	"todo-app/internal/model"
)

type AuthService struct {
	UserService *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{UserService: NewUserService()}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *AuthService) Authenticate(email, password string) (*model.User, error) {
	user, err := service.UserService.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
