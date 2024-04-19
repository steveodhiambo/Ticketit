package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/steveodhiambo/ticket-it/service/user"
	"log"
	"net/http"
)

type Server struct {
	addr string
	db   *sql.DB
}

// NewServer creates a new instance of the API server
func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on:", s.addr)
	return http.ListenAndServe(s.addr, router)
}
