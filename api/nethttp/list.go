package main

import (
	"log/slog"
	"net/http"
)

/*
Handle the error by returning HTTP 500 as status code, and error message as body.

Return the JSON result using the WriteJSONResult method!
*/

func (s *Server) ListPetsHandler(w http.ResponseWriter, r *http.Request) {
	pets, err := s.shelter.SelectAllPets()
	slog.Debug("SelectAllPets called", slog.Any("pets", pets), slog.Any("err", err))
}
