package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"restGolang/config"
	"restGolang/controller"
	"restGolang/model"
)

func main() {
	r := gin.Default()

	db, err := config.ConnectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Users{})

	controller.InitUserController(db)

	r.POST("/testing/post", controller.CreatePost)
	
    r.GET("/all/user", controller.AllUser)
	r.GET("/api/find", controller.FindUserBy)
	r.PUT("/user/edit", controller.EditUser)
	r.DELETE("/delete", controller.DeleteUser)
    
	r.Run(":8080")
}
