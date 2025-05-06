package service

import "todo-app/internal/model"

type TaskTagService interface {
	GetTaskTags() ([]model.Tag, error)
	CreateTaskTag(*model.Tag) error
	DeleteTaskTag(ID int) error
}
