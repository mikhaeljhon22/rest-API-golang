package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"restGolang/model"
	"restGolang/repository"
	"restGolang/service"
	"net/http"
)

var userService service.UserService
func InitUserController(db *gorm.DB) {
	// userRepo := repository.NewUserRepository(db)
	userRepo := repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
}

func CreatePost(c *gin.Context) {
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to insert data", "data": user})
}

func AllUser(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil || len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no data found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func EditUser(c *gin.Context) {
	var input struct {
		ID       uint   `json:"ID"`
		Username string `json:"Username"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userService.UpdateUser(input.ID, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to update"})
}

func DeleteUser(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username required"})
		return
	}

	err := userService.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to delete"})
}

func FindUserBy(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username required"})
		return
	}

	user, err := userService.FindUser(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
