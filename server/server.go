package server

import (
	"todolist/activity"
	"todolist/config"
	"todolist/todo"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	db := config.GetDBConnection()
	
	activityRepo := activity.NewActivityRepository(db)
	activityUsecase := activity.NewActivityUsecase(activityRepo)
	activityController := activity.NewActivityController(activityUsecase)

	activityRoutes(r, activityController)

	todoRepo := todo.NewTodoRepository(db)
	todoUsecase := todo.NewTodoUsecase(todoRepo)
	todoController := todo.NewTodoController(todoUsecase)

	todoRoutes(r, todoController)
}
