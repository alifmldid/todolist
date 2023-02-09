package main

import (
	"todolist/server"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	server.RegisterAPIService(r)

	r.Run(":3030")
}