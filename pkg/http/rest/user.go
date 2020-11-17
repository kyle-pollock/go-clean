package rest

import (
	"encoding/json"
	"github.com/kyle-pollock/go-clean/pkg/entities"
	"log"
	"net/http"
)

type UserService interface {
	GetAllUsers() ([]*entities.User, error)
}

func (rest *rest) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := rest.userService.GetAllUsers()
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
