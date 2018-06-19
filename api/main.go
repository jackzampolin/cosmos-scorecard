package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents the API server
type Server struct {
	Port int `json:"port"`

	Version string
	Commit  string
	Branch  string
}

// Router returns the router
func (s *Server) Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	return router
}
