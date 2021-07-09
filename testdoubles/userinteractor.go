package testdoubles

import "github.com/kyle-pollock/go-clean/usecases/user"

type UserInteractorDummy struct{}

func (s *UserInteractorDummy) GetAllUsers(presenter user.Presenter) error { return nil }
