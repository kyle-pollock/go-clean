package user

import (
	"sort"
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/entities"
)

func TestGetUser(t *testing.T) {
	if testing.Short() {
		t.Skip("db integration test")
	}

	db, err := NewMySQLTest()
	if err != nil {
		t.Fatal(err)
	}
	if err = Setup(db, "./setup_test.sql"); err != nil {
		t.Fatal(err)
	}
	defer Teardown(db, "./teardown_test.sql")

	gateway := New(db)

	t.Run("return users", func(t *testing.T) {
		wantUsers := []*entities.User{
			entities.NewUser(1, "testuser green"),
			entities.NewUser(2, "testuser foo"),
			entities.NewUser(3, "testuser bar"),
		}
		users, err := gateway.GetAllUsers()
		if err != nil {
			t.Errorf("failed to get users: %v", err)
		}

		sort.Slice(wantUsers, func(i, j int) bool { return wantUsers[i].ID() < wantUsers[j].ID() })
		sort.Slice(users, func(i, j int) bool { return users[i].ID() < users[j].ID() })

		for i := range wantUsers {
			if wantUsers[i].ID() != users[i].ID() {
				t.Errorf("wrong id. want <%d> got <%d>", wantUsers[i].ID(), users[i].ID())
			}
			if wantUsers[i].Name() != users[i].Name() {
				t.Errorf("wrong name. want <%s> got <%s>", wantUsers[i].Name(), users[i].Name())
			}
		}
	})
}
