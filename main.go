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
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	// Initialize Gin router
	r := gin.Default()


	//cors permission
	r.Use(cors.New(cors.Config {
		 AllowOrigins:     []string{"*"}, // izinkan semua origin
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
	}))

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

	master := r.Group("/api/service")
	source:= r.Group("/api/local", middleware.AuthGuard(userService))

	master.POST("/testing/post", controller.CreatePost)
	master.GET("/all/user", controller.AllUser)
	master.GET("/api/find", controller.FindUserBy)
	master.PUT("/user/edit", controller.EditUser)
	master.DELETE("/delete", controller.DeleteUser)
	master.POST("/create/acc", controller.CreateAcc)
	master.POST("/login", controller.Login)

	//uploaded file 
	master.POST("/upload", controller.SaveFileHandler)
	source.GET("/home", controller.Home)
	r.Run(":8080"); 
}
