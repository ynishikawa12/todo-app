package service

import (
	"errors"
	"todo-app/internal/model"
	"todo-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{UserRepository: repo}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *AuthServiceImpl) Authenticate(email, password string) (*model.User, error) {
	user, err := service.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
