package main 

import "github.com/gin-gonic/gin"

type User struct {
	Username string `json:"Username"`
	Gender string `json:"Gender"`
}

func setupRouter() *gin.Engine{
	router := gin.Default()
	user := User{
		Username: "Mikhael",
		Gender: "Male",
	}
	router.GET("/tes", func(c *gin.Context){
		c.JSON(200, user)
	})
	return router
}
func main(){
router := setupRouter()
router.Run(":8081")
}