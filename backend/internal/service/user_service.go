package service

import (
	"golang.org/x/crypto/bcrypt"
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{UserRepository: repository.NewUserRepository()}
}

func (service *UserService) GetUserByID(id int) (*model.User, error) {
	return service.UserRepository.GetUserByID(id)
}

func (service *UserService) CreateUser(user *model.User) (int, error) {
	return service.UserRepository.CreateUser(user)
}

func (service *UserService) GetUserByEmail(email string) (*model.User, error) {
	return service.UserRepository.GetUserByEmail(email)
}

func (service *UserService) UpdateUser(id int, updates model.User) error {
	return service.UserRepository.UpdateUser(id, updates)
}

func (service *UserService) GenerateHashedPassword(password string) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassowrd), nil
}
