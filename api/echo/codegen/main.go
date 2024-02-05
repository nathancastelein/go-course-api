package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-api/openapi"
)

func main() {
	server := NewServer()

	router := echo.New()
	router.Static("/swaggerui", "../../../openapi/swaggerui")

	openapi.RegisterHandlers(router, server)

	router.Logger.Fatal(router.Start(":8080"))
}
