package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveFiles)

	log.Println("Listening on :8844...")
	err := http.ListenAndServe(":8844", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./main.html"
	}
	http.ServeFile(w, r, p)
}
