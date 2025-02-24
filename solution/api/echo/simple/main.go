package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-api/solution/database/sqlboiler"
)

type Pet struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func ListPetsHandler(c echo.Context) error {
	pets, err := sqlboiler.SelectAllPets()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	jsonPets := make([]Pet, len(pets))
	for i, pet := range pets {
		jsonPets[i] = Pet{
			Id:       pet.Id,
			Name:     pet.Name,
			Category: pet.Category,
		}
	}

	return c.JSON(http.StatusOK, jsonPets)
}

func FindPetHandler(c echo.Context) error {
	// Get id from URL path
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "id must be an int")
	}

	pet, err := sqlboiler.SelectOnePet(idInt)
	if err != nil {
		return c.String(http.StatusNotFound, "pet not found")
	}

	return c.JSON(http.StatusOK, &Pet{
		Id:       pet.Id,
		Name:     pet.Name,
		Category: pet.Category,
	})
}

func main() {
	e := echo.New()

	e.GET("/pet", ListPetsHandler)
	e.GET("/pet/:id", FindPetHandler)

	e.Logger.Print("Server is running on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
