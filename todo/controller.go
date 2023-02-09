package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController interface{
	GetAllTodo(c *gin.Context)
	GetTodoById(c *gin.Context)
	CreateTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type todoController struct{
	todoUsecase TodoUsecase
}

func NewTodoController(todoUsecase TodoUsecase) TodoController{
	return &todoController{todoUsecase}
} 

func (controller *todoController) GetAllTodo(c *gin.Context){
	activity_group_id := c.Query("activity_group_id")
	todos, err := controller.todoUsecase.GetAllTodo(c, activity_group_id)
	
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "Success",
		"message": "Success",
		"data": todos,
	})
}

func (controller *todoController) GetTodoById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}

	todo, err := controller.todoUsecase.GetTodoById(c, id)
	
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "Success",
		"message": "Success",
		"data": todo,
	})
}

func (controller *todoController) CreateTodo(c *gin.Context){
	var payload Payload
	c.ShouldBindJSON(&payload)

	todo, err := controller.todoUsecase.CreateTodo(c, payload)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "Success",
		"message": "Success",
		"data": todo,
	})	
}

func (controller *todoController) UpdateTodo(c *gin.Context){
	var payload Payload

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}

	c.ShouldBindJSON(&payload)

	todo, err := controller.todoUsecase.UpdateTodo(c, id, payload)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"status": "Success",
		"message": "Success",
		"data": todo,
	})
}

func (controller *todoController) DeleteTodo(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": "Fail",
			"message": err.Error(),
			"data": "",
		})
		return
	}

	status, err := controller.todoUsecase.DeleteTodo(c, id)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, 
		gin.H{
			"status": status,
			"message": err.Error(),
			"data": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "Success",
		"message": "Success",
		"data": "",
	})
}