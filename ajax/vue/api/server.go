package api

import (
	"encoding/json"
	"log"
	"net/http"

	t "github.com/vue-examples/ajax/types"
	"github.com/vugu/vugu/devutil"
)

type Server struct {
	Router     *devutil.Mux
	statements *t.StatementList
}

func NewServer() *Server {
	server := &Server{
		Router:     devutil.NewMux(),
		statements: t.NewStatementList(),
	}

	server.initRoutes()
	server.statements.AddExampleData()

	return server
}

func (s *Server) initRoutes() {
	s.Router.Exact("/", s.serveVueFiles())
	s.Router.Exact("/static/styles.css", s.serveCSS())
	s.Router.Exact("/api/allStatements", s.getStatements())
	s.Router.Exact("/api/newStatement", s.submitQuestion())
	s.Router.Exact("/api/submitAnswer", s.submitAnswer())
}

func (s *Server) submitQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statement t.Statement
		err := json.NewDecoder(r.Body).Decode(&statement)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.statements.AddStatement(statement)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Println("Added new statement to the server...")

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(statement)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) submitAnswer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var answer struct {
			ID     int
			Answer int
		}

		err := json.NewDecoder(r.Body).Decode(&answer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.statements.AddToStatement(answer.ID, answer.Answer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		log.Printf("Server received answer\n-----\nQuestion ID: %d\nSelected answer: %d\n", answer.ID, answer.Answer)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(answer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s Server) getStatements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(s.statements.List)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s Server) serveVueFiles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := "." + r.URL.Path
		if p == "./" {
			p = "./index.html"
		}
		http.ServeFile(w, r, p)
	}
}

func (s Server) serveCSS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/styles.css")
	}
}
