package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tendermint/tendermint/rpc/client"
	// "github.com/tendermint/tendermint/rpc/core/types"
)

// Server represents the API server
type Server struct {
	Port int            `json:"port"`
	Gaia TendermintConn `json:"gaia"`

	Client *client.HTTP

	Version string
	Commit  string
	Branch  string
}

// TendermintConn is a representation of a  connection
type TendermintConn struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (g TendermintConn) tcp() string {
	return fmt.Sprintf("tcp://%s:%s", g.Host, g.Port)
}

// Router returns the router
func (s *Server) Router() *mux.Router {
	router := mux.NewRouter()
	s.Client = client.NewHTTP(s.Gaia.tcp(), "/websocket")
	router.HandleFunc("/hello", s.HealthHandler)
	router.HandleFunc("/status", s.StatusHandler)
	return router
}

// HealthHandler handles the /health route
func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("{\"status\": \"ok\"}"))
}

// StatusHandler handles the /health route
func (s *Server) StatusHandler(w http.ResponseWriter, r *http.Request) {
	stat, err := s.Client.DumpConsensusState()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	out := &RoundState{}
	err = json.Unmarshal(stat.RoundState, out)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	byt, err := json.Marshal(out.Validators)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(byt)
}
