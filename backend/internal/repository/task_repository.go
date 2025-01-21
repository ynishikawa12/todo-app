package repository

import (
	"todo-app/internal/model"
)

type TaskRepository interface {
	GetTaskByID(taskID int) (*model.Task, error)
	GetTasksByUserID(userID int) ([]model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(taskID int, updates model.Task) error
	DeleteTask(taskID int) error
}
