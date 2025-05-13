package repository

import (
	"crypto/sha256"
	"gorm.io/gorm"
	"restGolang/model"
	"encoding/hex"
	"errors"
)

type UserRepository interface {
	Create(user *model.Users) error 
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (*model.Users, error)
	UpdateUsername(id uint, username string) error
	DeleteByUsername(username string) error
	CreateAcc(user *model.UserNews) error
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
	pwHash := hashPassword(userNews.Password)
	userNews.Password = pwHash
	find := r.db.Where("username = ?", userNews.Username).First(&userNews)

	if(find.RowsAffected == 0){
      return r.db.Create(userNews).Error
	}else{
		return errors.New(`username already exist`)
	}
}
func hashPassword(password string)string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}