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

func main() {
	// Handle swagger UI
	fs := http.FileServer(http.Dir("../../../../api/swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	// Handler GET /pet
	http.HandleFunc("/pet", ListPetsHandler)

	log.Println("Server is running on http://localhost:8080")
	log.Println("Swagger UI URL: http://localhost:8080/swaggerui/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
