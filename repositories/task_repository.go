package repositories

import (
	"context"
	"task-manager/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type TaskRepository interface {
// 	Create(title string) models.Task
// 	GetAll() []models.Task
// 	Delete(id int) bool
// 	Update(id int, title, status string) (models.Task, bool)
// 	GetByID(id int) (models.Task, bool)
// 	Patch(id int, title *string, status *string) (models.Task, bool)
// }

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(connection *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: connection}
}

func (repository *TaskRepository) Create(ctx context.Context, task models.CreateTaskRequest) (int, error) {
	var id int

	err := repository.db.QueryRow(ctx, "insert into tasks(title) values ($1) returning id", task.Title).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

func (repository *TaskRepository) FindTasks(ctx context.Context) ([]models.Task, error) {
	sql := "select id, title, status from tasks order by id"

	rows, err := repository.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Status,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repository *TaskRepository) FindTaskById(ctx context.Context, id int) (models.Task, error) {
	sql := "select id, title, status from tasks t where t.id = $1"

	var task models.Task

	err := repository.db.QueryRow(ctx, sql, id).Scan(
		&task.ID,
		&task.Title,
		&task.Status,
	)

	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
