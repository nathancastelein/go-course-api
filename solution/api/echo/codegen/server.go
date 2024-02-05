package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-api/openapi"
	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/nathancastelein/go-course-api/solution/database/sqlboiler"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func newOpenAPIPet(pet shelter.Pet) openapi.Pet {
	petId := int64(pet.Id)
	return openapi.Pet{
		Id:       &petId,
		Name:     pet.Name,
		Category: pet.Category,
	}
}

// Lists Pets
// (GET /pet)
func (s *Server) ListPets(ctx echo.Context) error {
	pets, err := sqlboiler.SelectAllPets()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "internal server error")
	}

	jsonPets := make([]openapi.Pet, len(pets))
	for i, pet := range pets {
		jsonPets[i] = newOpenAPIPet(pet)
	}

	return ctx.JSON(http.StatusOK, jsonPets)
}

// Finds Pets by categories
// (GET /pet/findByCategories)
func (s *Server) FindPetsByCategories(ctx echo.Context, params openapi.FindPetsByCategoriesParams) error {
	pets, err := sqlboiler.SelectByCategories(params.Categories...)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "internal server error")
	}

	jsonPets := make([]openapi.Pet, len(pets))
	for i, pet := range pets {
		jsonPets[i] = newOpenAPIPet(pet)
	}

	return ctx.JSON(http.StatusOK, jsonPets)
}

// Find pet by ID
// (GET /pet/{petId})
func (s *Server) GetPetById(ctx echo.Context, petId int64) error {
	pet, err := sqlboiler.SelectOnePet(int(petId))
	if err != nil {
		return ctx.String(http.StatusNotFound, "pet not found")
	}

	return ctx.JSON(http.StatusOK, newOpenAPIPet(*pet))
}

// Add a new pet to the shelter
// (POST /pet)
func (s *Server) AddPet(ctx echo.Context) error {
	addPet := openapi.AddPetJSONRequestBody{}
	if err := ctx.Bind(&addPet); err != nil {
		return ctx.String(http.StatusBadRequest, "fail to read body")
	}

	pet, err := sqlboiler.InsertNewPet(&shelter.Pet{
		Name:     addPet.Name,
		Category: addPet.Category,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "fail to insert pet")
	}

	return ctx.JSON(http.StatusOK, newOpenAPIPet(*pet))
}

// Deletes a pet
// (DELETE /pet/{petId})
func (s *Server) DeletePet(ctx echo.Context, petId int64) error {
	err := sqlboiler.DeleteOnePet(int(petId))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "deletion failed")
	}
	return ctx.NoContent(http.StatusOK)
}

// Update an existing pet
// (PUT /pet/{petId})
func (s *Server) UpdatePet(ctx echo.Context, petId int64) error {
	updatePet := openapi.UpdatePetJSONRequestBody{}
	if err := ctx.Bind(&updatePet); err != nil {
		return ctx.String(http.StatusBadRequest, "fail to read body")
	}

	updatedPet := &shelter.Pet{
		Id:       int(petId),
		Name:     updatePet.Name,
		Category: updatePet.Category,
	}
	if err := sqlboiler.UpdatePet(updatedPet); err != nil {
		return ctx.String(http.StatusInternalServerError, "fail to update pet")
	}

	return ctx.JSON(http.StatusOK, newOpenAPIPet(*updatedPet))
}
