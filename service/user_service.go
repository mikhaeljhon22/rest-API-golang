package service

import (
	"fmt"
	"restGolang/model"
	"restGolang/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	// GORM
	CreateUser(user *model.Users) error
	GetAllUsers() ([]model.Users, error)
	FindUser(username string) (*model.Users, error)
	UpdateUser(id uint, username string) error
	DeleteUser(username string) error
	CreateAcc(userNews *model.UserNews) error
	Login(Username string, Password string) (*model.UserNews, error)

	// JWT
	GenerateJwt(data string) (string, error)
	VerifyJwt(tokenString string) error

	// Mongo
	CreateMongo(about *model.Mongos) error
}

type userService struct {
	repo       repository.UserRepository
	aboutRepo  repository.AboutRepository // MongoDB repo
}

func NewUserService(userRepo repository.UserRepository, aboutRepo repository.AboutRepository) UserService {
	return &userService{
		repo:      userRepo,
		aboutRepo: aboutRepo,
	}
}

var secretKey = []byte("secret-key")

func (s *userService) GenerateJwt(data string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": data,
			"exp":      time.Now().Add(time.Hour * 24 * 3).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *userService) VerifyJwt(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

// ------------------ GORM (PostgreSQL) ------------------

func (s *userService) CreateUser(user *model.Users) error {
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]model.Users, error) {
	return s.repo.FindAll()
}

func (s *userService) FindUser(username string) (*model.Users, error) {
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

func (s *userService) Login(username string, password string) (*model.UserNews, error) {
	return s.repo.Login(username, password)
}

func (s *userService) CreateMongo(about *model.Mongos) error {
	return s.aboutRepo.Create(about)
}
