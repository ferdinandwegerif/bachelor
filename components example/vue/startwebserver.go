package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveFiles)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe("127.0.0.1:3000", nil)
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
