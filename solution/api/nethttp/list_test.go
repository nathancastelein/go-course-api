package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListPetsHandler(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Arrange
		require := require.New(t)
		req := httptest.NewRequest(http.MethodGet, "/pet", nil)
		response := httptest.NewRecorder()
		server := NewServer(&petRepositoryStub{require: require})

		// Act
		server.ListPetsHandler(response, req)

		// Assert
		require.Equal(http.StatusOK, response.Result().StatusCode, "invalid HTTP status code received")
		body, err := io.ReadAll(response.Body)
		require.NoError(err, "fail to read response body")

		require.JSONEq(`[
			{
				"Id": 1,
				"Name": "Daffy",
				"Category": "duck"
			},
			{
				"Id": 2,
				"Name": "Mickey",
				"Category": "mouse"
			}
		]`, string(body), "invalid response body received")
	})

	t.Run("error", func(t *testing.T) {
		// Arrange
		require := require.New(t)
		req := httptest.NewRequest(http.MethodGet, "/pet", nil)
		response := httptest.NewRecorder()
		server := NewServer(&petRepositoryStubError{})

		// Act
		server.ListPetsHandler(response, req)

		// Assert
		require.Equal(http.StatusInternalServerError, response.Result().StatusCode, "invalid HTTP status code received")
		body, err := io.ReadAll(response.Body)
		require.NoError(err, "fail to read response body")
		require.Equal(`an error message`, string(body), "invalid response body received")
	})
}
