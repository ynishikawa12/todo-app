package handler

import (
	"net/http"
	"strconv"

	"todo-app/internal/consts"
	"todo-app/internal/model"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskTagHandler struct {
	TaskTagService service.TaskTagService
}

func NewTaskTagHandler(s service.TaskTagService) *TaskTagHandler {
	return &TaskTagHandler{TaskTagService: s}
}

func (handler *TaskTagHandler) GetTaskTags(c *gin.Context) {
	tasks, err := handler.TaskTagService.GetTaskTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (handler *TaskTagHandler) CreateTaskTag(c *gin.Context) {
	var taskTag model.Tag
	if err := c.ShouldBindJSON(&taskTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.TaskTagService.CreateTaskTag(&taskTag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, taskTag)
}

func (handler *TaskTagHandler) DeleteTaskTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(consts.TaskIDParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := handler.TaskTagService.DeleteTaskTag(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
