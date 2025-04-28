package main

import (
	"github.com/gin-gonic/gin"
	"restGolang/controller" 
)

func main() {
	r := gin.Default()

	r.GET("/testing/get", controller.GetUserStruct)
	r.Run(":8080")
}
