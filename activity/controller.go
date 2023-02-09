package activity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController interface{
	GetAllActivity(c *gin.Context)
	GetActivityById(c *gin.Context)
	CreateActivity(c *gin.Context)
	UpdateActivity(c *gin.Context)
	DeleteActivity(c *gin.Context)
}

type activityController struct{
	activityUsecase ActivityUsecase	
}

func NewActivityController(activityUsecase ActivityUsecase) ActivityController{
	return &activityController{activityUsecase}
} 

func (controller *activityController) GetAllActivity(c *gin.Context){
	activities, err := controller.activityUsecase.GetAllActivity(c)
	
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
		"data": activities,
	})
}

func (controller *activityController) GetActivityById(c *gin.Context){
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

	activity, err := controller.activityUsecase.GetActivityById(c, id)
	
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
		"data": activity,
	})
}

func (controller *activityController) CreateActivity(c *gin.Context){
	var payload Payload
	c.ShouldBindJSON(&payload)

	activity, err := controller.activityUsecase.CreateActivity(c, payload)
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
		"data": activity,
	})	
}

func (controller *activityController) UpdateActivity(c *gin.Context){
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

	activity, err := controller.activityUsecase.UpdateActivity(c, id, payload)

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
		"data": activity,
	})
}

func (controller *activityController) DeleteActivity(c *gin.Context){
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

	status, err := controller.activityUsecase.DeleteActivity(c, id)

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