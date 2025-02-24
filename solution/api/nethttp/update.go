package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type updatePet struct {
	NewName string `json:"newName"`
}

func (s *Server) UpdatePetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := PathInt(r, "id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid pet id %d", id)
		return
	}

	body, err := Body(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "fail to read body")
		return
	}

	updateInfos := updatePet{}
	if err := json.Unmarshal(body, &updateInfos); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid body: %s", err.Error())
		return
	}

	if updateInfos.NewName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing new name")
		return
	}

	err = s.shelter.UpdatePetName(id, updateInfos.NewName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
