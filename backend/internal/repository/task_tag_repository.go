package repository

import "todo-app/internal/model"

type TaskTagRepository interface {
	GetTaskTags() ([]model.Tag, error)
	CreateTaskTag(*model.Tag) error
	DeleteTaskTag(ID int) error
}
