package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kyle-pollock/go-clean/pkg/http/rest"
)

func main() {
	log.Print("starting process")

	host := flag.String("host", os.Getenv("GO_HOST"), "The host to listen on")
	port := flag.String("port", os.Getenv("GO_PORT"), "The port to listen on")
	flag.Parse()

	rest := rest.New(isReady)

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
