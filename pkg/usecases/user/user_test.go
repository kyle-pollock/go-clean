package user_test

import (
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func Test(t *testing.T) {
	userInteractor := user.New(&testdoubles.UserGatewayStub{})
	t.Run("get users", func(t *testing.T) {
		t.Parallel()
		users, err := userInteractor.GetAllUsers(context.Background())
		if err != nil {
			t.Error(err)
		}
		if len(users) == 0 {
			t.Error("user failed to read")
		}
	})
}
