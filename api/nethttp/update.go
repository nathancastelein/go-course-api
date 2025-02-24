package main

import (
	"log/slog"
	"net/http"
)

/*
Get ID from path parameter.

Read new name from JSON body, using the Body method.

Unmarshal the body into a dedicated struct or into a map, to get the new name.

Handle errors as usual.
*/

func (s *Server) UpdatePetHandler(w http.ResponseWriter, r *http.Request) {
	id := 42
	newName := "?"
	err := s.shelter.UpdatePetName(id, newName)
	slog.Debug("UpdatePetName called", slog.Any("err", err))
}
