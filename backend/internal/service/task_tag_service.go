package service

import "todo-app/internal/model"

type TaskTagService interface {
	GetTags() ([]model.TaskTag, error)
	CreateTag(*model.TaskTag) error
	DeleteTag(ID int) error
}
