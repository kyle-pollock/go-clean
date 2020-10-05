package rest_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
)

func TestGetUsers(t *testing.T) {
	isReadyFn := func() error { return nil }
	server := newServer(
		"/",
		isReadyFn,
		&testdoubles.UserServiceStub{},
	)
	expectedStatusCode := http.StatusOK

	t.Run("get users", func(t *testing.T) {
		t.Parallel()
		wantResponse := `[{"id":1,"name":"testuser greene"},{"id":2,"name":"testuser foo"},{"id":3,"name":"testuser bar"}]`
		request, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Error(err)
		}
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != expectedStatusCode {
			t.Errorf(errorFormat, response.Code, expectedStatusCode)
		}
		gotResponse := strings.TrimSpace(response.Body.String())
		if strings.Compare(gotResponse, wantResponse) != 0 {
			t.Errorf("want: %s\ngot: %s\n", wantResponse, gotResponse)
		}
	})
}
