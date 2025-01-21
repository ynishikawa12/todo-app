package repository

import (
	"todo-app/internal/database"
	"todo-app/internal/model"

	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskRepository() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{DB: database.DB}
}

func (repo *TaskRepositoryImpl) GetTaskByID(taskID int) (*model.Task, error) {
	var task model.Task
	if err := repo.DB.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (repo *TaskRepositoryImpl) GetTasksByUserID(userID int) ([]model.Task, error) {
	var tasks []model.Task
	if err := repo.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *TaskRepositoryImpl) CreateTask(task *model.Task) error {
	if err := repo.DB.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepositoryImpl) UpdateTask(ID int, updates model.Task) error {
	if err := repo.DB.Model(&model.Task{}).Where("id = ?", ID).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepositoryImpl) DeleteTask(ID int) error {
	if err := repo.DB.Delete(&model.Task{}, ID).Error; err != nil {
		return err
	}
	return nil
}
