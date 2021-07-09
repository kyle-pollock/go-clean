package entities

import (
	"testing"
)

func TestIsQualifyGreenMiles(t *testing.T) {
	t.Run("user does qualify for green miles", func(t *testing.T) {
		user := NewUser(1, "Greene User")
		if !user.IsQualifyGreenMiles() {
			t.Errorf("want true got false")
		}
	})

	t.Run("user does not qualify for green miles", func(t *testing.T) {
		user := NewUser(1, "Test User")
		if user.IsQualifyGreenMiles() {
			t.Errorf("want true got false")
		}
	})
}
