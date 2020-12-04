package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
)

func TestGetUsers(t *testing.T) {
	expectedStatusCode := http.StatusOK
	server := serverSetup(
		&testdoubles.UserInteractorDummy{},
	)

	t.Run("get users", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Error(err)
		}
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		if response.Code != expectedStatusCode {
			t.Errorf(errorFormat, response.Code, expectedStatusCode)
		}
	})
}
