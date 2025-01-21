package service

import (
	"todo-app/internal/model"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) (int, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(id int, updates model.User) error
}
