package user_test

import (
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func Test(t *testing.T) {

	presenter := new(testdoubles.UsersPresenterSpy)
	useCase := user.New(new(testdoubles.UserGatewayStub))

	err := useCase.GetAllUsers(presenter)
	if err != nil {
		t.Error(err)
	}

	if !presenter.PresentWasCalled {
		t.Error("usecase failed to call presenter")
	}

	viewModel := presenter.ViewModel

	if len(viewModel) == 0 {
		t.Error("usecase failed to read users")
	}
}
