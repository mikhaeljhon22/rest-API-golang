	package main 

	import (
		"github.com/gin-gonic/gin"
		"net/http"
	)
	type User struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
	func main(){
		router := gin.Default()

		router.GET("/tes", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "success",
				"status": http.StatusOK,
			})
		})

		router.POST("/post/tes", func (c *gin.Context){
			var user User 

			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Message": "invalid body request",
					"error": err.Error(),
				})
				return 
			}

			c.JSON(200, gin.H{
				"message": "success",
				"Data": user,
			})
		})
		
		router.GET("/get/struct", func (c *gin.Context){
			var user User 
			user.Username = "Mikhael"
			user.Password = "Akuanakkaya123"
		
			c.JSON(200,user)
		})

		router.Run(":8081")
	}