package json

import (
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func TestPresenter(t *testing.T) {
	presenter := new(Presenter)
	t.Run("validateViewModel", func(t *testing.T) {
		t.Parallel()
		want := UsersViewModel{
			Users: []*UserViewModel{
				{
					ID:   1,
					Name: "testUser1",
				},
				{
					ID:   2,
					Name: "testUser2",
				},
			},
		}

		presenter.Present([]*user.UserResponseModel{
			{
				ID:   1,
				Name: "testUser1",
			},
			{
				ID:   2,
				Name: "testUser2",
			},
		})

		for i, userViewModel := range presenter.ViewModel().Users {
			if *userViewModel != *want.Users[i] {
				t.Errorf("wrong presentation:\ngot:  %v\nwant %v", userViewModel, want.Users[i])
			}
		}
	})
}
