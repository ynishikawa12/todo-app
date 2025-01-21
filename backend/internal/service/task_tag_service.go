package service

import "todo-app/internal/model"

type TaskTagService interface {
	GetTaskTags() ([]model.TaskTag, error)
	CreateTaskTag(*model.TaskTag) error
	DeleteTaskTag(ID int) error
}
