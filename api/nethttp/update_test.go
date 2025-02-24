package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdatePetName(t *testing.T) {
	// Arrange
	require := require.New(t)
	req := httptest.NewRequest(http.MethodGet, "/pet/1", bytes.NewBufferString(`{"newName": "Daisy"}`))
	req.SetPathValue("id", "1")
	response := httptest.NewRecorder()
	spy := &petRepositorySpy{petRepositoryStub: petRepositoryStub{require: require}}
	server := NewServer(spy)

	// Act
	server.UpdatePetHandler(response, req)

	// Assert
	require.True(spy.updateCalled, "UpdatePetName should be called")
	require.Equal(http.StatusOK, response.Result().StatusCode, "invalid HTTP status code received")
	body, err := io.ReadAll(response.Body)
	require.NoError(err, "fail to read response body")
	require.Empty(body, "invalid response body received")
}
