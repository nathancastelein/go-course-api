package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListOnePet(t *testing.T) {
	// Arrange
	require := require.New(t)
	req := httptest.NewRequest(http.MethodGet, "/pet/1", nil)
	req.SetPathValue("id", "1")
	response := httptest.NewRecorder()
	server := NewServer(&petRepositoryStub{require: require})

	// Act
	server.GetOnePetHandler(response, req)

	// Assert
	require.Equal(http.StatusOK, response.Result().StatusCode, "invalid HTTP status code received")
	body, err := io.ReadAll(response.Body)
	require.NoError(err, "fail to read response body")

	require.JSONEq(`{
		"Id": 1,
		"Name": "Daffy",
		"Category": "duck"
	}`, string(body), "invalid response body received")
}

func TestListOnePet_Lowercase(t *testing.T) {
	// Arrange
	require := require.New(t)
	req := httptest.NewRequest(http.MethodGet, "/pet/1?format=lowercase", nil)
	req.SetPathValue("id", "1")
	response := httptest.NewRecorder()
	server := NewServer(&petRepositoryStub{require: require})

	// Act
	server.GetOnePetHandler(response, req)

	// Assert
	require.Equal(http.StatusOK, response.Result().StatusCode, "invalid HTTP status code received")
	body, err := io.ReadAll(response.Body)
	require.NoError(err, "fail to read response body")

	require.JSONEq(`{
		"id": 1,
		"name": "Daffy",
		"category": "duck"
	}`, string(body), "invalid response body received")
}
