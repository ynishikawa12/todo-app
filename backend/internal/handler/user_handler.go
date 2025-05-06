package handler

import (
	"net/http"
	"strconv"

	"todo-app/internal/consts"
	"todo-app/internal/model"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{UserService: s}
}

// GetUserByID はユーザーIDからユーザー情報を取得する
// @Summary      ユーザーを返す
// @Description  ユーザーIDから取得する
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   query      int  true  "ユーザーID"
// @Success      200  {object}   model.User
// @Failure      400  {object}   model.ErrorResponse
// @Router       /users [get]
func (handler *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(consts.UserIDParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid user ID"})
		return
	}

	user, err := handler.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser はユーザーを作成する
// @Summary      ユーザーを作成する
// @Description  ユーザーを作成する
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user   body   model.User  true  "ユーザー情報"
// @Success      201  {object}   int "ユーザーID"
// @Failure      400  {object}   model.ErrorResponse
// @Router       /users [post]
func (handler *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := handler.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(consts.UserIDParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.UserService.UpdateUser(id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
