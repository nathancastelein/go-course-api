package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryParamString(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/?input=bar", nil)

	// Act
	param := QueryParamString(req, "input")

	// Assert
	require.Equal(t, "bar", param)
}

func TestQueryParamInt(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/?input=42", nil)

	// Act
	param, err := QueryParamInt(req, "input")

	// Assert
	require.Nil(t, err)
	require.Equal(t, 42, param)
}

func TestHeader(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Content-Type", "application/json")

	// Act
	param := Header(req, "Content-Type")

	// Assert
	require.Equal(t, "application/json", param)
}

func TestPathParameter(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetPathValue("name", "Ada")

	// Act
	param := Path(req, "name")

	// Assert
	require.Equal(t, "Ada", param)
}

func TestPathParameterInt(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetPathValue("id", "42")

	// Act
	param, err := PathInt(req, "id")

	// Assert
	require.NoError(t, err)
	require.Equal(t, 42, param)
}

func TestBody(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(`Hello World`))

	// Act
	param, err := Body(req)

	// Assert
	require.NoError(t, err)
	require.Equal(t, "Hello World", string(param))
}
