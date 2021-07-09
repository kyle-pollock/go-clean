package user

import (
	"github.com/kyle-pollock/go-clean/entities"
)

type Interactor interface {
	GetAllUsers(Presenter) error
}

type Presenter interface {
	Present([]*UserResponseModel)
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

func (s *service) GetAllUsers(presenter Presenter) error {
	users, err := s.gateway.GetAllUsers()
	if err != nil {
		return err
	}

	var responseModel []*UserResponseModel
	for _, user := range users {
		responseModel = append(responseModel, mapUserEntityToResponseModel(user))
	}
	presenter.Present(responseModel)

	return nil
}

func mapUserEntityToResponseModel(user *entities.User) *UserResponseModel {
	return &UserResponseModel{
		ID:   user.ID(),
		Name: user.Name(),
	}
}
