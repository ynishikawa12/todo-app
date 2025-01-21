package repository

import (
	"todo-app/internal/model"
)

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (int, error)
	UpdateUser(id int, updates model.User) error
}
