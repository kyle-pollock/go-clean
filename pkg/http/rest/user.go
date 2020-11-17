package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

func (rest *rest) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := rest.userInteractor.GetAllUsers()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
