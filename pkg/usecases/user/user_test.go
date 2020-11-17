package user

import (
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
)

func Test(t *testing.T) {
	repo := &testdoubles.UserGatewayStub{}

	t.Run("get users", func(t *testing.T) {
		t.Parallel()
		userInteractor := New(repo)
		users, err := userInteractor.GetAllUsers()
		if err != nil {
			t.Error(err)
		}
		if len(users) == 0 {
			t.Error("user failed to read")
		}
	})
}
