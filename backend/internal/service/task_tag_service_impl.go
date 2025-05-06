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

func (service *TaskTagServiceImpl) GetTaskTags() ([]model.Tag, error) {
	return service.TaskTagRepository.GetTaskTags()
}

func (service *TaskTagServiceImpl) CreateTaskTag(taskTag *model.Tag) error {
	return service.TaskTagRepository.CreateTaskTag(taskTag)
}

func (service *TaskTagServiceImpl) DeleteTaskTag(ID int) error {
	return service.TaskTagRepository.DeleteTaskTag(ID)
}
