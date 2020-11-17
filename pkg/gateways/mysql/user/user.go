package user

import (
	"database/sql"
	"fmt"

	"github.com/kyle-pollock/go-clean/pkg/entities"
)

type userGateway struct {
	db *sql.DB
}

func New(db *sql.DB) *userGateway {
	return &userGateway{db: db}
}

func (gateway *userGateway) GetAllUsers() ([]*entities.User, error) {
	dtos, err := gateway.getDTOs()
	if err != nil {
		return nil, err
	}
	return userMapper(dtos), nil
}

func (gateway *userGateway) getDTOs() ([]*userDTO, error) {
	rows, err := gateway.db.Query(`
		SELECT id, name
		FROM goclean.user;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to read from mysql: %v", err)
	}

	dtos := []*userDTO{}
	for rows.Next() {
		dto := new(userDTO)

		if err := rows.Scan(&dto.id, &dto.name); err != nil {
			return nil, fmt.Errorf("failed to mysql scan a user: %v", err)
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
