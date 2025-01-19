package service

import (
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
