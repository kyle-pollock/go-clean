package rest

import (
	"log"
	"net/http"

	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

type rest struct {
	isReady        func() error
	userInteractor user.Interactor
}

func New(isReady func() error, userInteractor user.Interactor) *rest {
	return &rest{
		isReady:        isReady,
		userInteractor: userInteractor,
	}
}

func (rest *rest) NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/readyz", http.HandlerFunc(rest.readyz))
	mux.Handle("/livez", http.HandlerFunc(rest.livez))
	mux.Handle("/users", http.HandlerFunc(rest.getUsers))
	return mux
}

func (rest *rest) readyz(w http.ResponseWriter, r *http.Request) {
	if err := rest.isReady(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	w.WriteHeader(http.StatusOK)
}

func (rest *rest) livez(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
