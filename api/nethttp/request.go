package main

import (
	"net/http"
)

func QueryParamString(request *http.Request, paramName string) string {
	return ""
}

func QueryParamInt(request *http.Request, paramName string) (int, error) {
	return 0, nil
}

func Header(request *http.Request, headerName string) string {
	return ""
}

func Path(request *http.Request, parameterName string) string {
	return ""
}

func PathInt(request *http.Request, parameterName string) (int, error) {
	return 0, nil
}

func Body(request *http.Request) ([]byte, error) {
	return nil, nil
}
