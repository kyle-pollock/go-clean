package userservice

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type Repository interface {
	GetAllUsers() ([]*entities.User, error)
}

type userService struct {
	repository Repository
}

func New(r Repository) *userService {
	return &userService{repository: r}
}

func (s *userService) GetAllUsers() ([]*entities.User, error) {
	return s.repository.GetAllUsers()
}
