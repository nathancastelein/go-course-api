package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathancastelein/go-course-api/database/sqlboiler"
)

func ListPetsHandler(w http.ResponseWriter, r *http.Request) {
	pets, err := sqlboiler.SelectAllPets()
	// If error, write response header HTTP 500 then return
	if err != nil {
		fmt.Fprint(w, "an error occured")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("found %d pets", len(pets)) // Remove me

	// Tips: have a look on encoding/json package!
	// https://pkg.go.dev/encoding/json
	// And JSON and Go article: https://go.dev/blog/json

	// Write response header HTTP 200
	w.WriteHeader(http.StatusOK)
	// Write result on response
	fmt.Fprint(w, "Hello world")
}

func AddNewPetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the body from the request
	// Tips:
	// Look at the Body field on http.Request
	// And look at the io.ReadAll method from io package: https://pkg.go.dev/io#ReadAll

	// Define a structure to unmarshal the body
	// Unmarshal the body with JSON function: https://pkg.go.dev/encoding/json#Unmarshal

	// Add the new pet using sqlboiler.InsertNewPet and get the inserted pet

	// Marshal the result and write it on response
}

func main() {
	// Handle swagger UI
	fs := http.FileServer(http.Dir("../../openapi/swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	// Handler GET and POST /pet
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
