package repositories

import "task-manager/models"

type TaskRepository interface {
	Create(title string) models.Task
	GetAll() []models.Task
	Delete(id int) bool
	Update(id int, title, status string) (models.Task, bool)
	GetByID(id int) (models.Task, bool)
	Patch(id int, title *string, status *string) (models.Task, bool)
}
