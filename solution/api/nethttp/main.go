package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nathancastelein/go-course-api/shelter"
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

type NewPet struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func AddNewPetHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "fail to read body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newPet := &NewPet{}
	if err := json.Unmarshal(body, newPet); err != nil {
		fmt.Fprint(w, "fail to unmarshal body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pet, err := sqlboiler.InsertNewPet(shelter.Pet{
		Name:     newPet.Name,
		Category: newPet.Category,
	})
	if err != nil {
		fmt.Fprint(w, "an error occured")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonPetsBytes, err := json.MarshalIndent(Pet{
		Name:     pet.Name,
		Category: pet.Category,
		Id:       pet.Id,
	}, "", "    ")
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

func main() {
	// Handle swagger UI
	fs := http.FileServer(http.Dir("../../../openapi/swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	// Handler GET /pet
	http.HandleFunc("/pet", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListPetsHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			AddNewPetHandler(w, r)
			return
		}
	})

	log.Println("Server is running on http://localhost:8080")
	log.Println("Swagger UI URL: http://localhost:8080/swaggerui/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
