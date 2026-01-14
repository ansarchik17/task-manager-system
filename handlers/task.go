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

// func (handler *TaskHandler) PostTask(c *gin.Context) {
// 	var request models.CreateTaskRequest

// 	err := c.BindJSON(&request)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind json body"))
// 		return
// 	}

// 	if request.Title == "" {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("title is required"))
// 		return
// 	}

// 	task, err := handler.service.Create(request.Title)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.NewApiError(err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, task)
// }

// func (handler *TaskHandler) GetTask(c *gin.Context) {
// 	tasks := handler.service.GetTasks()
// 	c.JSON(http.StatusOK, tasks)
// }

// func (handler *TaskHandler) GetTaskById(c *gin.Context) {
// 	idStr := c.Param("id")

// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid id"))
// 		return
// 	}

// 	task, err := handler.service
// }

// func (handler *TaskHandler) DeleteTask(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid id"))
// 		return
// 	}

// 	err = handler.service.DeleteTask(id)

// 	if err != nil {
// 		c.JSON(http.StatusNotFound, models.NewApiError("Task does not found"))
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }

// func (handler *TaskHandler) PutTask(c *gin.Context) {

// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("invalid id"))
// 		return
// 	}

// 	// 2. JSON body
// 	var req models.UpdateTaskDTO
// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("invalid json body"))
// 		return
// 	}

// 	// 3. PUT validation (қатаң)
// 	if req.Title == "" || req.Status == "" {
// 		c.JSON(http.StatusBadRequest, models.NewApiError("title and status are required"))
// 		return
// 	}

// 	// 4. Service
// 	updatedTask, err := handler.service.UpdateTask(id, req.Title, req.Status)
// 	if err != nil {
// 		if err.Error() == "task not found" {
// 			c.JSON(http.StatusNotFound, models.NewApiError(err.Error()))
// 			return
// 		}
// 		c.JSON(http.StatusBadRequest, models.NewApiError(err.Error()))
// 		return
// 	}
// 	// 5. Response
// 	c.JSON(http.StatusOK, updatedTask)
// }

// func (hander *TaskHandler) PatchTask(c *gin.Context) {

// }
