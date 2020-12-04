package testdoubles

import "github.com/kyle-pollock/go-clean/pkg/usecases/user"

type UserInteractorDummy struct{}

func (s *UserInteractorDummy) GetAllUsers(presenter user.Presenter) error { return nil }
