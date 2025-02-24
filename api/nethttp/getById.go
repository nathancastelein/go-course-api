package main

import (
	"log/slog"
	"net/http"
)

/*
Get the ID from the path value "id".
If you have an error, return an HTTP 400.

Then use SelectOnePet with the provided ID.
If you have an error, return an HTTP 500 with the error message.

Finally, write the JSON result!
*/

func (s *Server) GetOnePetHandler(w http.ResponseWriter, r *http.Request) {
	id := 42
	pet, err := s.shelter.SelectOnePet(id)
	slog.Debug("SelectOnePet called", slog.Any("pet", pet), slog.Any("err", err))
}
