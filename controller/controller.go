package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"restGolang/model"
	"restGolang/repository"
	"restGolang/service"
	"net/http"
	"fmt"
	"path/filepath"
	"github.com/google/uuid"
	"restGolang/dto"
	"restGolang/util"
	 "crypto/rand"
	 qrcode "github.com/skip2/go-qrcode"
)

//init service file
var userService service.UserService
func Init(db *gorm.DB, aboutRepo repository.AboutRepository) {
	repo := repository.NewUserRepository(db)
	userService = service.NewUserService(repo, aboutRepo)
}

func CreatePost(c *gin.Context) {
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userService.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
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

	user, err:= userService.FindUser(username)
	if err != nil {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateAcc(c *gin.Context){
	var userNews model.UserNews 

	if err:= c.BindJSON(&userNews); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result := userService.CreateAcc(&userNews)
	if result != nil {
		c.JSON(500,gin.H{
			"message": result.Error(),
		})
	}else{
	token, err:= userService.GenerateJwt(userNews.Username)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return 
		}

		//set header 
		fmt.Println(token)
	 c.Header("Authorization", "Bearer " + token)
	c.JSON(200, gin.H{
		"message:" : "succcess to create account",
	})
}
}
func Login(c *gin.Context){
	var input dto.LoginDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := userService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	token,err:= userService.GenerateJwt(input.Username)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Header("Authorization","Bearer "+token)
	fmt.Println(token)


	c.JSON(http.StatusOK, gin.H{
		"message": "success to login",
		"user": user,
	})
}

func Home(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "success to home",
	})
}

// upload file
func SaveFileHandler (c *gin.Context){
	file,err := c.FormFile("file") //create formFile

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "no file received",
		})
		return
	}
	//create file name
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension 


	if extension != ".jpg" {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "file must be jpg",
		})
		return
	}

	//save file
	if err := c.SaveUploadedFile(file, "uploadedFile/" + newFileName); err != nil{
		c.AbortWithStatusJSON(500, gin.H{
			"error": "unable to save",
		})
		fmt.Println("error")

	}

	c.JSON(200, gin.H{
		"message": "success upload file",
	})
}
func TestMongo(c *gin.Context){
	var input dto.MongosDTO
	if err:= c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := model.Mongos{
		Name:  input.Name,
		Email: input.Email,
	}
	if err := userService.CreateMongo(&data); err != nil {
		c.JSON(500, gin.H{"error": "failed insert to mongo"})
		return
	}

	c.JSON(200, gin.H{"message": "inserted"})

}

func SendMail(c *gin.Context){
	name := c.Query("target")
	result:= util.SendMail(name)
	if result != nil {
		c.JSON(500, gin.H{
		"error": result,
	})
	}else{
	c.JSON(200, gin.H{
		"message": "success to send",
	})
}
}
func UUID(c *gin.Context){
	newUUID,err := uuid.NewRandom()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	c.JSON(200,gin.H{
		"message": newUUID,
	})
}

func RandomNumb(c *gin.Context){
    p, _ := rand.Prime(rand.Reader, 15)
	c.JSON(200,gin.H{
		"random": p,
	})
}

func QRGenerator(c *gin.Context){
	 var png []byte
  png, err := qrcode.Encode("https://google.com", qrcode.Medium, 256)

  if err != nil {
	fmt.Println(err)
  }
  c.Writer.Write(png)
}