package rest

import (
	"encoding/json"
	"log"
	"net/http"

	presenter "github.com/kyle-pollock/go-clean/pkg/presenters/json"
)

func (rest *rest) getUsers(w http.ResponseWriter, r *http.Request) {
	presenter := new(presenter.Presenter)

	err := rest.userInteractor.GetAllUsers()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(presenter.ViewModel()); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
