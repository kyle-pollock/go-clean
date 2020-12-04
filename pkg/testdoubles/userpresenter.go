package testdoubles

import (
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

type UsersPresenterSpy struct {
	PresentWasCalled bool
	ViewModel        []*user.UserResponseModel
}

func (p *UsersPresenterSpy) Present(model []*user.UserResponseModel) {
	p.PresentWasCalled = true
	p.ViewModel = model
}
