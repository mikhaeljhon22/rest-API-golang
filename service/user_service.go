package service

import (
	"restGolang/model"
	"restGolang/repository"
)

type UserService interface {
	CreateUser(user *model.Users) error
	GetAllUsers() ([]model.Users, error)
	FindUser(username string) (*model.Users, error)
	UpdateUser(id uint, username string) error
	DeleteUser(username string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *model.Users) error {
	err := s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil 
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
