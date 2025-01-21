package repository

import "todo-app/internal/model"

type TaskTagRepository interface {
	GetTaskTags() ([]model.TaskTag, error)
	CreateTaskTag(*model.TaskTag) error
	DeleteTaskTag(ID int) error
}
