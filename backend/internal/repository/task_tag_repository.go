package repository

import "todo-app/internal/model"

type TaskTagRepository interface {
	GetTags() ([]model.TaskTag, error)
	CreateTag(*model.TaskTag) error
	DeleteTag(ID int) error
}

