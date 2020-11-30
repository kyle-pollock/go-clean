package user

import (
	"context"

	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type Interactor interface {
	GetAllUsers(context.Context, Presenter) error
}

type Gateway interface {
	GetAllUsers() ([]*entities.User, error)
}

type Presenter interface {
	Present([]*UserResponseModel)
}

type service struct {
	gateway Gateway
}

func New(r Gateway) Interactor {
	return &service{gateway: r}
}

func (s *service) GetAllUsers(ctx context.Context, presenter Presenter) error {
	users, err := s.gateway.GetAllUsers()
	if err != nil {
		return err
	}

	var response []*UserResponseModel
	for _, user := range users {
		response = append(response, mapUserEntityToResponseModel(user))
	}

	presenter.Present(response)

	return nil
}

func mapUserEntityToResponseModel(user *entities.User) *UserResponseModel {
	return &UserResponseModel{
		ID:   user.ID(),
		Name: user.Name(),
	}
}
