package controller

import (
	"github.com/gin-gonic/gin"
	 "gorm.io/gorm"
	"restGolang/model"
)

func CreatePost(c *gin.Context, db *gorm.DB){
		var user model.Users
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return 
		}

		result := db.Create(&user)
		if result.Error != nil{
			c.JSON(500,gin.H{
				"message": result.Error.Error(),
			})
		}
		c.JSON(200, gin.H{
			"message": "success to insert data",
			"data": user,
		})
}