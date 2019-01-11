package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Route Route //负责api模块
}

type Route interface {
	Endpoint(*mux.Router)
}

func (s *Server) Run() error {
	r := mux.NewRouter()
	s.Route.Endpoint(r)
	port := ":8200"
	srv := &http.Server{
		Handler: r,
		Addr:    port,
	}
	log.Printf("server run on %s", port)
	return srv.ListenAndServe()
}
