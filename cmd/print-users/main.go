package main

import (
	"fmt"
	"log"

	mysql "github.com/kyle-pollock/go-clean/pkg/storage/mysql/user"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func main() {
	db, err := mysql.NewMySQLTest()
	if err != nil {
		log.Fatalf("Error in initializing database: %v", err)
	}
	userRepo := mysql.New(db)
	userService := user.New(userRepo)

	users, err := userService.GetAllUsers()

	for _, user := range users {
		fmt.Print(user)
	}
}
