package user

import (
	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type userDTO struct {
	id   int
	name string
}

func userMapper(dtos []*userDTO) []*entities.User {
	var users []*entities.User
	for _, dto := range dtos {
		users = append(users, entities.NewUser(
			dto.id,
			dto.name,
		))
	}
	return users
}
