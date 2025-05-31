package repository

import (
	"gorm.io/gorm"
	"restGolang/model"
	"restGolang/util"
	"errors"
	"fmt"
)

type UserRepository interface {
	Create(user *model.Users) error 
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (*model.Users, error)
	UpdateUsername(id uint, username string) error
	DeleteByUsername(username string) error
	CreateAcc(user *model.UserNews) error
	Login(Username string, Password string)(*model.UserNews, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.Users) error{
	return r.db.Create(&user).Error
}

func (r *userRepository) FindAll() ([] model.Users, error){
	var user []model.Users
	result := r.db.Find(&user).Error
	return user,result
}
func (r *userRepository) FindByUsername(username string) (*model.Users, error){
	var user model.Users 
	result := r.db.First(&user,"username = ?", username).Error
	return &user,result
}

func (r *userRepository) UpdateUsername(id uint, username string) error {
	return r.db.Model(&model.Users{}).Where("id = ?", id).Update("Username", username).Error
}

func (r *userRepository) DeleteByUsername(username string) error {
	return r.db.Where("username = ?", username).Delete(&model.Users{}).Error
}
func (r *userRepository) CreateAcc(userNews *model.UserNews) error{
	pwHash := util.HashPassword(userNews.Password)
	userNews.Password = pwHash
	find := r.db.Where("username = ? OR email = ?", userNews.Username, userNews.Email).First(&userNews)

	if(find.RowsAffected == 0){
      return r.db.Create(userNews).Error
	}else{
		return errors.New(`username or email already exist`)
	}
}


func (r *userRepository) Login(Username string, Password string) (*model.UserNews, error) {
	hashPw := util.HashPassword(Password)
	var user model.UserNews
	result := r.db.Where("username = ? AND password = ?", Username, hashPw).First(&user)
	
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return &user, nil
}


