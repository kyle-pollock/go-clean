package main

import (
	"fmt"
	"log"

	mysql "github.com/kyle-pollock/go-clean/pkg/gateways/mysql/user"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func main() {
	db, err := mysql.NewMySQLTest()
	if err != nil {
		log.Fatalf("error in initializing database: %v", err)
	}
	userInteractor := user.New(mysql.New(db))

	users, err := userInteractor.GetAllUsers()
	if err != nil {
		log.Fatalf("error getting all users: %v", err)
	}

	for _, user := range users {
		fmt.Print(user)
	}
}
