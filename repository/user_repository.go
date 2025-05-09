package repository

// init gorm and go model	
import (
	"gorm.io/gorm"
	"restGolang/model"
)

type UserRepository interface {
	Create(user *model.Users) error
	FindAll() ([]model.Users, error)
	FindByUsername(username string) (*model.Users, error)
	UpdateUsername(id uint, username string) error
	DeleteByUsername(username string) error
}

type userRepository struct {
	db *gorm.DB
}

//create dependency injection userRepository with pointer gorm DB
func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{db}
}
func (r *userRepository) Create(user *model.Users) error {
	return r.db.Create(&user).Error
}


func (r *userRepository) FindAll() ([]model.Users, error) {
	var users []model.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByUsername(username string) (*model.Users, error) {
	var user model.Users
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUsername(id uint, username string) error {
	return r.db.Model(&model.Users{}).Where("id = ?", id).Update("Username", username).Error
}

func (r *userRepository) DeleteByUsername(username string) error {
	return r.db.Where("username = ?", username).Delete(&model.Users{}).Error
}

