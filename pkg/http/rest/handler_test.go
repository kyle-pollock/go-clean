package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyle-pollock/go-clean/pkg/testdoubles"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"

	"github.com/kyle-pollock/go-clean/pkg/http/rest"
)

const (
	errorFormat = "Got status code '%d' but expected '%d'"
)

func serverSetup(userInteractor user.Interactor) *http.ServeMux {
	isReadyFn := func() error { return nil }
	server := http.NewServeMux()
	server.Handle("/", rest.New(
		isReadyFn,
		userInteractor,
	).NewHandler())
	return server
}

func TestRest(t *testing.T) {
	server := serverSetup(
		&testdoubles.UserInteractorStub{},
	)
	expectedStatusCode := http.StatusOK

	t.Run("healthz", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequest(http.MethodGet, "/readyz", nil)
		if err != nil {
			t.Error(err)
		}
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != expectedStatusCode {
			t.Errorf(errorFormat, response.Code, expectedStatusCode)
		}
	})

	t.Run("livez", func(t *testing.T) {
		t.Parallel()
		request, err := http.NewRequest(http.MethodGet, "/livez", nil)
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
