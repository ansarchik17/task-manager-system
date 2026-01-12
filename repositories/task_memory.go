package repositories

import (
	"task-manager/models"
)

type InMemoryTaskRepository struct {
	tasks  map[int]models.Task
	nextID int
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (r *InMemoryTaskRepository) Create(title string) models.Task {
	task := models.Task{
		ID:     r.nextID,
		Title:  title,
		Status: "new",
	}
	r.tasks[r.nextID] = task
	r.nextID++
	return task
}

func (r *InMemoryTaskRepository) GetAll() []models.Task {
	result := []models.Task{}

	for _, task := range r.tasks {
		result = append(result, task)
	}

	return result
}

func (r *InMemoryTaskRepository) Delete(id int) bool {
	if _, exists := r.tasks[id]; !exists {
		return false
	}
	delete(r.tasks, id)
	return true
}

func (r *InMemoryTaskRepository) Update(id int, status, title string) (models.Task, bool) {
	task, exists := r.tasks[id]

	if !exists {
		return models.Task{}, false
	}

	task.Title = title
	task.Status = status

	r.tasks[id] = task
	return task, true
}

func (r *InMemoryTaskRepository) GetByID(id int) (models.Task, bool) {
	task, exists := r.tasks[id]
	if !exists {
		return models.Task{}, false
	}
	return task, true
}

func (r *InMemoryTaskRepository) Patch(id int, title *string, status *string) (models.Task, bool) {
	task, exists := r.tasks[id]
	if !exists {
		return models.Task{}, false
	}

	if title != nil {
		task.Title = *title
	}
	if status != nil {
		task.Status = *status
	}

	r.tasks[id] = task
	return task, true
}
