package testdoubles

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type UserRepositoryStub struct{}

func (r *UserRepositoryStub) GetAllUsers() ([]*entities.User, error) {
	return []*entities.User{
		entities.NewUser(1, "testuser green"),
		entities.NewUser(2, "testuser foo"),
		entities.NewUser(3, "testuser bar"),
	}, nil
}
