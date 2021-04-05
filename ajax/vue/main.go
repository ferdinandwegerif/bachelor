package main

import (
	"log"
	"net/http"

	"github.com/vue-examples/ajax/api"
)

func main() {
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)
	server := api.NewServer()
	log.Fatal(http.ListenAndServe(l, server.Router))

}
