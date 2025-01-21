package service

import (
	"golang.org/x/crypto/bcrypt"
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: repo}
}

func (service *UserServiceImpl) GetUserByID(id int) (*model.User, error) {
	return service.UserRepository.GetUserByID(id)
}

func (service *UserServiceImpl) CreateUser(user *model.User) (int, error) {
	hashedPassword, err := generateHashedPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword
	return service.UserRepository.CreateUser(user)
}

func (service *UserServiceImpl) GetUserByEmail(email string) (*model.User, error) {
	return service.UserRepository.GetUserByEmail(email)
}

func (service *UserServiceImpl) UpdateUser(id int, updates model.User) error {
	return service.UserRepository.UpdateUser(id, updates)
}

func generateHashedPassword(password string) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassowrd), nil
}
