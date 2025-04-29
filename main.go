package main

import (
	"github.com/gin-gonic/gin"
	"restGolang/controller" 
	"restGolang/config"
    "restGolang/model"
    "log"
)

func main() {
	r := gin.Default()

	db, err := config.ConnectToPostgreSQL()
    if err != nil {
        log.Fatal(err)
    }

    err = db.AutoMigrate(&model.Users{})
    if err != nil {
        log.Fatal(err)
    }

    r.POST("/testing/post", func(c *gin.Context){
        controller.CreatePost(c, db)
    })
	
	r.Run(":8080")
}
