package controller

import (
	"github.com/gin-gonic/gin"
	"restGolang/model"
)

func GetUserStruct(c *gin.Context) {
	user := model.User{
		Username: "Mikhael",
		Password: "Akuanakkaya123",
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}
