package models

type PatchTaskRequest struct {
	Title  *string `json:"title"`
	Status *string `json:"status"`
}
