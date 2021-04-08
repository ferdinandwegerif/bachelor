package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	t "github.com/vugu-examples/ajax/types"
	"github.com/vugu/vugu/devutil"
)

type Server struct {
	Router     *devutil.Mux
	statements *t.StatementList
}

//Insert empty string to only run as REST API without generating wasm front-end files. For testing purposes of the API
func NewServer(dir string) *Server {
	server := &Server{
		Router:     devutil.NewMux(),
		statements: t.NewStatementList(),
	}
	server.initRoutes()
	if dir != "" {
		server.initWasm(dir)
	}
	server.statements.AddExampleData()

	return server
}

func (s *Server) initRoutes() {
	s.Router.Exact("/api/allStatements", s.getStatements())
	s.Router.Exact("/api/newStatement", s.submitQuestion())
	s.Router.Exact("/api/submitAnswer", s.submitAnswer())
	/* 	s.HandleFunc("/api/allStatements", s.getStatements()).Methods("GET")
	   	s.HandleFunc("/api/getStatement/{id}", s.getStatement()).Methods("GET")
	   	s.HandleFunc("/api/newStatement", s.submitQuestion()).Methods("POST")
	   	s.HandleFunc("/api/submitAnswer", s.submitAnswerToQuestion()).Methods("POST") */
	//s.HandleFunc("api/getAnswers", s.)

	/* s.Exact("/api/allStatements", http.Handler(s.getStatements())) */
}

func (s *Server) initWasm(dir string) {
	fmt.Println("initializing webassembly files")
	wc := devutil.NewWasmCompiler().SetDir(dir)
	s.Router.Match(devutil.NoFileExt, devutil.DefaultAutoReloadIndex.Replace(
		`<!-- styles -->`,
		`<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
		<link rel="stylesheet" href="/static/styles.css">
	`))
	s.Router.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	s.Router.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	s.Router.Default(devutil.NewFileServer().SetDir(dir))
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

/* func (s Server) getStatement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idstr)
		fmt.Println("Statement id that is requested:", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		statement, err := s.statements.GetStatementByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(statement)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
} */
