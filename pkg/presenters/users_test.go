package presenters

import (
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func TestUsersPresenter(t *testing.T) {
	wantViewModel := ViewModel{
		Users: []*User{
			{ID: 1, Name: "testuser greene"},
			{ID: 2, Name: "testuser foo"},
			{ID: 3, Name: "testuser bar"},
		},
	}

	presenter := new(UsersPresenter)
	presenter.Present([]*user.UserResponseModel{
		{ID: 1, Name: "testuser greene"},
		{ID: 2, Name: "testuser foo"},
		{ID: 3, Name: "testuser bar"},
	})

	for i, user := range presenter.ViewModel().Users {
		if user.ID != wantViewModel.Users[i].ID {
			t.Errorf("\ngot  %d\nwant %d", user.ID, wantViewModel.Users[i].ID)
		}
		if user.Name != wantViewModel.Users[i].Name {
			t.Errorf("\ngot  %s\nwant %s", user.Name, wantViewModel.Users[i].Name)
		}
	}
}
