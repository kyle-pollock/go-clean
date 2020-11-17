package testdoubles

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type UserGatewayStub struct{}

func (r *UserGatewayStub) GetAllUsers() ([]*entities.User, error) {
	return []*entities.User{
		entities.NewUser(1, "testuser greene"),
		entities.NewUser(2, "testuser foo"),
		entities.NewUser(3, "testuser bar"),
	}, nil
}
