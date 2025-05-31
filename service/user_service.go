package service

import (
	"restGolang/model"
	"restGolang/repository"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
)

type UserService interface {
	CreateUser(user *model.Users) error
	GetAllUsers() ([]model.Users, error)
	FindUser(username string) (*model.Users, error)
	UpdateUser(id uint, username string) error
	DeleteUser(username string) error
	CreateAcc(userNews *model.UserNews) error
	Login(Username string, Password string) (*model.UserNews, error)
	GenerateJwt(data string) (string, error)
	VerifyJwt(tokenString string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService{
	return &userService{repo}
}

var secretKey = []byte("secret-key")

func (s *userService) GenerateJwt(data string) (string, error){
token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": data, 
        "exp": time.Now().Add(time.Hour * 24 * 3).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func (s *userService) VerifyJwt(tokenString string) error {
token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
	return secretKey, nil
})
if err != nil {
	return err
}

if !token.Valid {
	return fmt.Errorf("invalid token")
}
return nil
}



func (s *userService) CreateUser(user *model.Users) error{
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]model.Users, error){
	return s.repo.FindAll()
}

func (s *userService) FindUser(username string) (*model.Users, error){
	return s.repo.FindByUsername(username)
}

func (s *userService) UpdateUser(id uint, username string) error {
	return s.repo.UpdateUsername(id, username)
}

func (s *userService) DeleteUser(username string) error {
	return s.repo.DeleteByUsername(username)
}


func (s *userService) CreateAcc(userNews *model.UserNews) error {
	return s.repo.CreateAcc(userNews)
}

func (s *userService) Login(Username string, Password string)(*model.UserNews, error){
	return s.repo.Login(Username,Password)
}
