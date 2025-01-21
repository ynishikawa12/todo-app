package service

import "todo-app/internal/model"

type TaskService interface {
	GetTasksByUserID(userID int) ([]model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(taskID int, updates model.Task) (*model.Task, error)
	DeleteTask(taskID int) error
}
