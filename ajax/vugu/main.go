// +build !wasm

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vugu-examples/ajax/api"
)

func main() {
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	server := api.NewServer(wd)
	log.Fatal(http.ListenAndServe(l, server.Router))
}
