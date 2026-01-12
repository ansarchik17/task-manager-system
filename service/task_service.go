package service

import (
	"errors"
	"task-manager/models"
	"task-manager/repositories"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (service *TaskService) Create(title string) (models.Task, error) {
	if title == "" {
		return models.Task{}, errors.New("title cannot be empty")
	}
	return service.repo.Create(title), nil
}

func (service *TaskService) GetTasks() []models.Task {
	return service.repo.GetAll()
}

func (service *TaskService) DeleteTask(id int) error {
	if ok := service.repo.Delete(id); !ok {
		return errors.New("task not found")
	}
	return nil
}

func (service *TaskService) UpdateTask(id int, title string, status string) (models.Task, error) {
	if title == "" {
		return models.Task{}, errors.New("title cannot be empty")
	}
	if status == "" {
		return models.Task{}, errors.New("status cannot be empty")
	}

	updatedTask, ok := service.repo.Update(id, title, status)
	if !ok {
		return models.Task{}, errors.New("task not found")
	}
	return updatedTask, nil
}

func (service *TaskService) PatchTask(id int, title *string, status *string) (models.Task, error) {
	if title == nil && status == nil {
		return models.Task{}, errors.New("no fields to update")
	}

	updatedTask, ok := service.repo.Patch(id, title, status)

	if !ok {
		return models.Task{}, errors.New("task not found")
	}

	return updatedTask, nil
}
