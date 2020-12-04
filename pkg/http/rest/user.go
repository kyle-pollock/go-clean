package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kyle-pollock/go-clean/pkg/presenters"
)

func (rest *rest) getUsers(w http.ResponseWriter, r *http.Request) {

	presenter := new(presenters.UsersPresenter)
	err := rest.userInteractor.GetAllUsers(presenter)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	viewModel := presenter.ViewModel()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(viewModel); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
