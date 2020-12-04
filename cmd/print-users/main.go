package main

import (
	"fmt"
	"log"

	mysql "github.com/kyle-pollock/go-clean/pkg/gateways/mysql/user"
	"github.com/kyle-pollock/go-clean/pkg/presenters"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

func main() {
	db, err := mysql.NewMySQLTest()
	if err != nil {
		log.Fatalf("error in initializing database: %v", err)
	}
	userInteractor := user.New(mysql.New(db))

	presenter := new(presenters.UsersPresenter)
	err = userInteractor.GetAllUsers(presenter)
	if err != nil {
		log.Fatalf("error getting all users: %v", err)
	}

	for _, user := range presenter.ViewModel().Users {
		fmt.Print(user)
	}
}
