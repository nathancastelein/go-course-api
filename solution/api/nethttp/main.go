package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nathancastelein/go-course-api/solution/database/sqlboiler"
)

type Pet struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func ListPetsHandler(w http.ResponseWriter, r *http.Request) {
	pets, err := sqlboiler.SelectAllPets()
	// If error, write response header HTTP 500 then return
	if err != nil {
		fmt.Fprint(w, "an error occured")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonPets := make([]Pet, len(pets))
	for i, pet := range pets {
		jsonPets[i] = Pet{
			Id:       pet.Id,
			Name:     pet.Name,
			Category: pet.Category,
		}
	}

	// encode to json
	jsonPetsBytes, err := json.MarshalIndent(jsonPets, "", "    ")
	if err != nil {
		fmt.Fprint(w, "an error occured")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response header HTTP 200
	w.WriteHeader(http.StatusOK)
	// Write result on response
	fmt.Fprint(w, string(jsonPetsBytes))
}
