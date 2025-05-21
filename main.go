package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"restGolang/config"
	"restGolang/controller"
	"restGolang/model"
	"restGolang/middleware"
	"restGolang/service"
	"restGolang/repository"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Connect to PostgreSQL database
	db, err := config.ConnectToPostgreSQL()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	} else {
		log.Println("Connected to PostgreSQL successfully")
	}

	 db.AutoMigrate(&model.Users{}, &model.UserNews{})


	controller.Init(db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	source:= r.Group("/api/local", middleware.AuthGuard(userService))

	r.POST("/testing/post", controller.CreatePost)
	r.GET("/all/user", controller.AllUser)
	r.GET("/api/find", controller.FindUserBy)
	r.PUT("/user/edit", controller.EditUser)
	r.DELETE("/delete", controller.DeleteUser)
	r.POST("/create/acc", controller.CreateAcc)
	r.POST("/login", controller.Login)
	source.GET("/home", controller.Home)
	r.Run(":8080"); 
}
