package repository

import (
	"todo-app/internal/database"
	"todo-app/internal/model"

	"gorm.io/gorm"
)

type TaskTagRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskTagRepository() *TaskTagRepositoryImpl {
	return &TaskTagRepositoryImpl{DB: database.DB}
}

func (repo *TaskRepositoryImpl) GetTags() ([]model.TaskTag, error) {
	var tags []model.TaskTag
	if err := repo.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *TaskRepositoryImpl) CreateTag(task *model.TaskTag) error {
	if err := repo.DB.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepositoryImpl) DeleteTaskTag(ID int) error {
	if err := repo.DB.Delete(&model.Task{}, ID).Error; err != nil {
		return err
	}
	return nil
}
