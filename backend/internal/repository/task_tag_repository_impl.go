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

func (repo *TaskTagRepositoryImpl) GetTaskTags() ([]model.Tag, error) {
	var tags []model.Tag
	if err := repo.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *TaskTagRepositoryImpl) CreateTaskTag(taskTag *model.Tag) error {
	if err := repo.DB.Create(&taskTag).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TaskTagRepositoryImpl) DeleteTaskTag(ID int) error {
	if err := repo.DB.Delete(&model.Task{}, ID).Error; err != nil {
		return err
	}
	return nil
}
