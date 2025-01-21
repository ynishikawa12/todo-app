package routes

import (
	"todo-app/internal/consts"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	taskTagRepository := repository.NewTaskTagRepository()
	taskTagService := service.NewTaskTagSerivce(taskTagRepository)
	taskTagHandler := handler.NewTaskTagHandler(taskTagService)

	r.POST(consts.UserURL, userHandler.CreateUser)
	r.GET(consts.UserURLWithID, userHandler.GetUserByID)
	r.PUT(consts.UserURLWithID, userHandler.UpdateUser)

	r.POST("/login", authHandler.Login)

	r.POST(consts.UserURLWithID+consts.TaskURL, taskHandler.CreateTask)
	r.GET(consts.UserURLWithID+consts.TaskURL, taskHandler.GetTasksByUserID)
	r.PUT(consts.UserURLWithID+consts.TaskURLWithID, taskHandler.UpdateTask)
	r.DELETE(consts.UserURLWithID+consts.TaskURLWithID, taskHandler.DeleteTask)

	r.GET(consts.TaskTagURL, taskTagHandler.GetTaskTags)
	r.POST(consts.TaskTagURL, taskTagHandler.CreateTaskTag)
	r.DELETE(consts.TaskTagURLWithID, taskTagHandler.DeleteTaskTag)

	return r
}
