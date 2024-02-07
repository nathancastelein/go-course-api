package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-api/openapi"
	"github.com/nathancastelein/go-course-api/shelter"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

/*
newOpenAPIPet creates an openapi.Pet from a shelter.Pet.
*/
func newOpenAPIPet(pet shelter.Pet) openapi.Pet {
	return openapi.Pet{
		Id:       pet.Id,
		Name:     pet.Name,
		Category: pet.Category,
	}
}

// Lists Pets
// (GET /pet)
func (s *Server) ListPets(ctx echo.Context) error {
	return nil
}

// Finds Pets by categories
// (GET /pet/findByCategories)
func (s *Server) FindPetsByCategories(ctx echo.Context, params openapi.FindPetsByCategoriesParams) error {
	return nil
}

// Find pet by ID
// (GET /pet/{petId})
func (s *Server) GetPetById(ctx echo.Context, petId int) error {
	return nil
}

// Add a new pet to the shelter
// (POST /pet)
func (s *Server) AddPet(ctx echo.Context) error {
	// Tips:
	// addPet := openapi.AddPetJSONRequestBody{}
	// https://echo.labstack.com/docs/binding
	// And Bind method from Context:
	// https://pkg.go.dev/github.com/labstack/echo/v4#Context
	return nil
}

// Deletes a pet
// (DELETE /pet/{petId})
func (s *Server) DeletePet(ctx echo.Context, petId int) error {
	return nil
}

// Update an existing pet
// (PUT /pet/{petId})
func (s *Server) RenamePetById(ctx echo.Context, petId int) error {
	// Tips:
	// updatePet := openapi.RenamePetByIdJSONBody{}
	// https://echo.labstack.com/docs/binding
	// And Bind method from Context:
	// https://pkg.go.dev/github.com/labstack/echo/v4#Context
	return nil
}
