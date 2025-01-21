package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{TaskRepository: repo}
}

func (s *TaskServiceImpl) GetTasksByUserID(userID int) ([]model.Task, error) {
	return s.TaskRepository.GetTasksByUserID(userID)
}

func (s *TaskServiceImpl) CreateTask(task *model.Task) error {
	return s.TaskRepository.CreateTask(task)
}

func (s *TaskServiceImpl) UpdateTask(taskID int, updates model.Task) (*model.Task, error) {
	if err := s.TaskRepository.UpdateTask(taskID, updates); err != nil {
		return nil, err
	}

	task, err := s.TaskRepository.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskServiceImpl) DeleteTask(taskID int) error {
	return s.TaskRepository.DeleteTask(taskID)
}
