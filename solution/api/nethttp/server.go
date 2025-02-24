package main

import (
	"log/slog"
	"net/http"

	"github.com/nathancastelein/go-course-api/shelter"
)

type Server struct {
	shelter shelter.PetRepository
	mux     *http.ServeMux
}

func NewServer(shelter shelter.PetRepository) *Server {
	server := &Server{
		shelter: shelter,
		mux:     http.NewServeMux(),
	}

	server.mux.HandleFunc("GET /pet", server.ListPetsHandler)
	server.mux.HandleFunc("GET /pet/{id}", server.GetOnePetHandler)
	server.mux.HandleFunc("PUT /pet/{id}", server.UpdatePetHandler)

	return server
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}
	slog.Info("Listening...", slog.String("address", "localhost:8080"))
	return server.ListenAndServe()
}
