package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	mysql "github.com/kyle-pollock/go-clean/pkg/gateways/mysql/user"
	"github.com/kyle-pollock/go-clean/pkg/http/rest"
	"github.com/kyle-pollock/go-clean/pkg/usecases/user"
)

var (
	host = os.Getenv("GO_HOST")
	port = os.Getenv("GO_PORT")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	flag.StringVar(&host, "host", host, "override default hostname")

	if port == "" {
		port = "3000"
	}
	flag.StringVar(&port, "port", port, "override default port")
}

func main() {
	log.Print("starting process")

	flag.Parse()

	db, err := mysql.NewMySQLTest()
	if err != nil {
		log.Fatalf("error in initializing database: %v", err)
	}
	userInteractor := user.New(mysql.New(db))

	rest := rest.New(isReady, userInteractor)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: rest.NewHandler(),
	}

	idleConnsClosed := make(chan struct{})
	go trapTerminationSignal(server, idleConnsClosed)

	log.Printf("services listening on %s:%s\n", host, port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	<-idleConnsClosed

	log.Print("end of processing")
}
