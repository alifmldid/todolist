package server

import (
	"todolist/activity"
	"todolist/todo"

	"github.com/gin-gonic/gin"
)

func activityRoutes(r *gin.Engine, controller activity.ActivityController){
	activity := r.Group("/activity-groups")
	activity.GET("/", controller.GetAllActivity)
	activity.GET("/:id", controller.GetActivityById)
	activity.POST("/", controller.CreateActivity)
	activity.PATCH("/:id", controller.UpdateActivity)
	activity.DELETE("/:id", controller.DeleteActivity)
}

func todoRoutes(r *gin.Engine, controller todo.TodoController){
	todo := r.Group("/todo-items")
	todo.GET("/", controller.GetAllTodo)
	todo.GET("/:id", controller.GetTodoById)
	todo.POST("/", controller.CreateTodo)
	todo.PATCH("/:id", controller.UpdateTodo)
	todo.DELETE("/:id", controller.DeleteTodo)
}