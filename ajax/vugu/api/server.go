package api

import (
	"encoding/json"
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
	s.Router.Exact("/api/allStatements", http.Handler(s.getStatements()))
}

func (s *Server) initWasm(dir string) {
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
		statement.ID = len(s.statements.List)
		s.statements.AddStatement(statement)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(statement)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) submitAnswerToQuestion() http.HandlerFunc {
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

func (s Server) getStatement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		err := json.NewDecoder(r.Body).Decode(&id)
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
}
