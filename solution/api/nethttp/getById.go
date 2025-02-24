package main

import (
	"fmt"
	"net/http"
)

func (s *Server) GetOnePetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := PathInt(r, "id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid pet id %d", id)
		return
	}

	pet, err := s.shelter.SelectOnePet(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	if QueryParamString(r, "format") == "lowercase" {
		type tmp struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Category string `json:"category"`
		}

		p := tmp{
			Id:       pet.Id,
			Name:     pet.Name,
			Category: pet.Category,
		}

		WriteJSONResult(w, http.StatusOK, p)
		return
	} else {
		WriteJSONResult(w, http.StatusOK, pet)
		return
	}
}
