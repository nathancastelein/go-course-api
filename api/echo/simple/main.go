package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-api/solution/database/sqlboiler"
)

type Pet struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func ListPetsHandler(c echo.Context) error {
	// Get all pets
	_, _ = sqlboiler.SelectAllPets()

	// Return http.InternalServerError with "internal server error" string in case of error
	// Return JSON
	// Tip: https://echo.labstack.com/docs/response#send-json
	return c.String(http.StatusOK, "Hello, World!")
}

func FindPetHandler(c echo.Context) error {
	// Get id from URL path
	// Return http.InternalServerError with "internal server error" string in case of error
	// Return JSON
	// Tips: https://echo.labstack.com/docs/response#send-json
	// Tips: https://echo.labstack.com/docs/quick-start#path-parameters
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.Static("/swaggerui", "../../../openapi/swaggerui")

	// Don't forget to register your routes
	// Tip: https://echo.labstack.com/docs/quick-start#routing

	e.Logger.Print("Server is running on http://localhost:8080")
	e.Logger.Print("Swagger UI URL: http://localhost:8080/swaggerui/")
	e.Logger.Fatal(e.Start(":8080"))
}
