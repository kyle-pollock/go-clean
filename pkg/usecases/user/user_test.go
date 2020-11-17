package user

import (
	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
	"testing"
)

func Test(t *testing.T) {
	repo := &testdoubles.UserGatewayStub{}

	t.Run("CRUD ops", func(t *testing.T) {
		t.Run("get users", func(t *testing.T) {
			t.Parallel()
			userservice := New(repo)
			users, err := userservice.GetAllUsers()
			if err != nil {
				t.Error(err)
			}
			if len(users) == 0 {
				t.Error("user failed to read")
			}
		})
	})
}
