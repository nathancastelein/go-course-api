package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintJSONResult(t *testing.T) {
	t.Run("test marshal ok", func(t *testing.T) {
		// Arrange
		response := httptest.NewRecorder()
		expectedResult := `{
	"foo": "bar"
}`
		expectedCode := http.StatusOK

		// Act
		WriteJSONResult(response, http.StatusOK, map[string]string{
			"foo": "bar",
		})

		// Assert
		require.Equal(t, expectedCode, response.Result().StatusCode)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		body, err := io.ReadAll(response.Body)
		require.NoError(t, err)
		require.JSONEq(t, expectedResult, string(body))
	})

	t.Run("test marshal error", func(t *testing.T) {
		// Arrange
		response := httptest.NewRecorder()
		expectedResult := http.StatusText(http.StatusInternalServerError)
		expectedCode := http.StatusInternalServerError

		// Act
		WriteJSONResult(response, http.StatusOK, func() {})

		// Assert
		require.Equal(t, expectedCode, response.Result().StatusCode)
		body, err := io.ReadAll(response.Body)
		require.NoError(t, err)
		require.Equal(t, expectedResult, string(body))
	})
}
