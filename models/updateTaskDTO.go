package models

type UpdateTaskDTO struct {
	Title  string `json: "title"`
	Status string `json: "status"`
}
