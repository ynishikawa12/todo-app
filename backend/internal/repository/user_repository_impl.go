package repository

import (
	"todo-app/internal/database"
	"todo-app/internal/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: database.DB}
}

func (repo *UserRepositoryImpl) GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) CreateUser(user *model.User) (int, error) {
	if err := repo.DB.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *UserRepositoryImpl) UpdateUser(id int, updates model.User) error {
	if err := repo.DB.Model(&model.User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}
