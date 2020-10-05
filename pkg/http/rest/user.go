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

type UserJSONDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (rest *rest) getUsers(w http.ResponseWriter, r *http.Request) {

	users, err := rest.userService.GetAllUsers()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dto := mapUsersJSONDTO(users)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}

func mapUsersJSONDTO(users []*entities.User) []*UserJSONDTO {
	usersDTO := []*UserJSONDTO{}
	for _, user := range users {
		usersDTO = append(usersDTO, &UserJSONDTO{
			ID:   user.ID(),
			Name: user.Name(),
		})
	}
	return usersDTO
}
