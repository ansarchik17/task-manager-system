package handler

import (
	"net/http"
	"strconv"
	"task-manager/models"
	"task-manager/repositories"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskHandler(taskRepo *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepo: taskRepo}
}

func (handler *TaskHandler) CreateTask(c *gin.Context) {
	var request models.CreateTaskRequest

	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("could not bind json body"))
		return
	}

	task := models.CreateTaskRequest{
		Title: request.Title,
	}

	taskId, err := handler.taskRepo.Create(c, task)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not create task"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": taskId})
}

func (handler *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := handler.taskRepo.FindTasks(c)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, tasks)
}

func (handler *TaskHandler) GetTaskById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid id"))
		return
	}
	task, err := handler.taskRepo.FindTaskById(c, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, task)
}

func (handler *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid id"))
		return
	}

	_, err = handler.taskRepo.FindTaskById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err.Error()))
		return
	}

	var request models.UpdateTaskDTO
	err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind json body"))
		return
	}

	updatedTask := models.Task{
		Title:  request.Title,
		Status: request.Status,
	}

	err = handler.taskRepo.Update(c, id, updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

func (handler *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid id"))
		return
	}

	err = handler.taskRepo.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
