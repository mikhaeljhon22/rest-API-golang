package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"restGolang/model"
)

func GetUserStruct(c *gin.Context) {
	user := model.User{
		Username: "Mikhael",
		Password: "Akuanakkaya123",
	}
	c.JSON(http.StatusOK, user)
}
