// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Pet defines model for Pet.
type Pet struct {
	Category string `json:"category"`
	Id       *int64 `json:"id,omitempty"`
	Name     string `json:"name"`
}

// FindPetsByCategoriesParams defines parameters for FindPetsByCategories.
type FindPetsByCategoriesParams struct {
	// Categories Categories values that need to be considered for filter
	Categories []string `form:"categories" json:"categories"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = Pet

// UpdatePetJSONRequestBody defines body for UpdatePet for application/json ContentType.
type UpdatePetJSONRequestBody = Pet

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Lists Pets
	// (GET /pet)
	ListPets(ctx echo.Context) error
	// Add a new pet to the shelter
	// (POST /pet)
	AddPet(ctx echo.Context) error
	// Finds Pets by categories
	// (GET /pet/findByCategories)
	FindPetsByCategories(ctx echo.Context, params FindPetsByCategoriesParams) error
	// Deletes a pet
	// (DELETE /pet/{petId})
	DeletePet(ctx echo.Context, petId int64) error
	// Find pet by ID
	// (GET /pet/{petId})
	GetPetById(ctx echo.Context, petId int64) error
	// Update an existing pet
	// (PUT /pet/{petId})
	UpdatePet(ctx echo.Context, petId int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListPets converts echo context to params.
func (w *ServerInterfaceWrapper) ListPets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListPets(ctx)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// FindPetsByCategories converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetsByCategories(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByCategoriesParams
	// ------------- Required query parameter "categories" -------------

	err = runtime.BindQueryParameter("form", true, true, "categories", ctx.QueryParams(), &params.Categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter categories: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FindPetsByCategories(ctx, params)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", ctx.Param("petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePet(ctx, petId)
	return err
}

// GetPetById converts echo context to params.
func (w *ServerInterfaceWrapper) GetPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", ctx.Param("petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPetById(ctx, petId)
	return err
}

// UpdatePet converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", ctx.Param("petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdatePet(ctx, petId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pet", wrapper.ListPets)
	router.POST(baseURL+"/pet", wrapper.AddPet)
	router.GET(baseURL+"/pet/findByCategories", wrapper.FindPetsByCategories)
	router.DELETE(baseURL+"/pet/:petId", wrapper.DeletePet)
	router.GET(baseURL+"/pet/:petId", wrapper.GetPetById)
	router.PUT(baseURL+"/pet/:petId", wrapper.UpdatePet)

}