package user

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type Repository interface {
	GetAllUsers() ([]*entities.User, error)
}

type UserService struct {
	repository Repository
}

func New(r Repository) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	return s.repository.GetAllUsers()
}
