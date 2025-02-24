package main

import (
	"fmt"
	"net/http"
)

func (s *Server) ListPetsHandler(w http.ResponseWriter, r *http.Request) {
	pets, err := s.shelter.SelectAllPets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	WriteJSONResult(w, http.StatusOK, pets)
}
