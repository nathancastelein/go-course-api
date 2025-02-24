package main

import (
	"io"
	"net/http"
	"strconv"
)

func QueryParamString(request *http.Request, paramName string) string {
	return request.URL.Query().Get(paramName)
}

func QueryParamInt(request *http.Request, paramName string) (int, error) {
	paramAsString := request.URL.Query().Get(paramName)
	return strconv.Atoi(paramAsString)
}

func Header(request *http.Request, headerName string) string {
	return request.Header.Get(headerName)
}

func Path(request *http.Request, parameterName string) string {
	return request.PathValue(parameterName)
}

func PathInt(request *http.Request, parameterName string) (int, error) {
	paramAsString := request.PathValue(parameterName)
	return strconv.Atoi(paramAsString)
}

func Body(request *http.Request) ([]byte, error) {
	body, err := io.ReadAll(request.Body)
	return body, err
}
