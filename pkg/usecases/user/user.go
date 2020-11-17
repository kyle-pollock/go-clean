package user

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type Interactor interface {
	GetAllUsers() ([]*UserResponseModel, error)
}

type Gateway interface {
	GetAllUsers() ([]*entities.User, error)
}

type service struct {
	gateway Gateway
}

func New(r Gateway) Interactor {
	return &service{gateway: r}
}

func (s *service) GetAllUsers() ([]*UserResponseModel, error) {
	users, err := s.gateway.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var response []*UserResponseModel
	for _, user := range users {
		response = append(response, mapUserEntityToResponseModel(user))
	}
	return response, nil
}

func mapUserEntityToResponseModel(user *entities.User) *UserResponseModel {
	return &UserResponseModel{
		ID: user.ID(),
		Name: user.Name(),
	}
}
