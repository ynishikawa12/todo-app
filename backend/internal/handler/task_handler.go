package handler

import (
	"net/http"
	"strconv"
	"todo-app/internal/consts"
	"todo-app/internal/model"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService service.TaskService
}

func NewTaskHandler(s service.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: s}
}

func (handler *TaskHandler) GetTaskByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(consts.UserIDParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": "invalid user ID"})
		return
	}

	tasks, err := handler.TaskService.GetTasksByUserID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}

	c.JSON(http.StatusOK, tasks)
}

func (handler *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.TaskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (handler *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(consts.TaskIDParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var updates model.Task
	if err := c.ShouldBindJSON((&updates)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := handler.TaskService.UpdateTask(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (handler *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(consts.TaskIDParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err != handler.TaskService.DeleteTask(id) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
