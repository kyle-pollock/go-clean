package user

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type Gateway interface {
	GetAllUsers() ([]*entities.User, error)
}

type UserService struct {
	repository Gateway
}

func New(r Gateway) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	return s.repository.GetAllUsers()
}
