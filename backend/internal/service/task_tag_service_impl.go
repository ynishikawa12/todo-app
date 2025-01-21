package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type TaskTagServiceImpl struct {
	TaskTagRepository repository.TaskTagRepository
}

func NewTaskTagSerivce(repo repository.TaskTagRepository) *TaskTagServiceImpl {
	return &TaskTagServiceImpl{TaskTagRepository: repo}
}

func GetTags() ([]model.TaskTag, error)
func CreateTag(*model.TaskTag) error
func DeleteTag(ID int) error
