package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kyle-pollock/go-clean/pkg/http/rest"
	mysql "github.com/kyle-pollock/go-clean/pkg/storage/mysql/user"
	"github.com/kyle-pollock/go-clean/pkg/userservice"
)

func main() {
	log.Print("starting process")

	host := flag.String("host", os.Getenv("GO_HOST"), "The host to listen on")
	port := flag.String("port", os.Getenv("GO_PORT"), "The port to listen on")
	flag.Parse()

	db, err := mysql.NewMySQLTest()
	if err != nil {
		log.Fatalf("Error in initializing database: %v", err)
	}
	userRepo := mysql.New(db)
	userService := userservice.New(userRepo)

	rest := rest.New(isReady, userService)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", *host, *port),
		Handler: rest.NewHandler(),
	}

	idleConnsClosed := make(chan struct{})
	go trapTerminationSignal(server, idleConnsClosed)

	log.Printf("services listening on %s:%s\n", *host, *port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	<-idleConnsClosed

	log.Print("end of processing")
}
