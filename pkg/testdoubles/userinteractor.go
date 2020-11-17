package testdoubles

import (
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

type UserInteractorStub struct{}

func (s *UserInteractorStub) GetAllUsers() ([]*user.UserResponseModel, error) {
	return []*user.UserResponseModel{
		{ID: 1, Name: "testuser greene"},
		{ID: 2, Name: "testuser foo"},
		{ID: 3, Name: "testuser bar"},
	}, nil
}
