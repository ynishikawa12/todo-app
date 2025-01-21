package service

import (
	"todo-app/internal/model"
)

type AuthService interface {
	Authenticate(email, password string) (*model.User, error)
}
