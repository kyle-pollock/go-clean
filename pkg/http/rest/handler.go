package rest

import (
	"log"
	"net/http"
)

type rest struct {
	isReady func() error
}

func New(isReady func() error) *rest {
	return &rest{
		isReady: isReady,
	}
}

func (rest *rest) NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/readyz", http.HandlerFunc(rest.readyz))
	mux.Handle("/livez", http.HandlerFunc(rest.livez))
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
