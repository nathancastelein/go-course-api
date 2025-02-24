package main

import (
	"net/http"
)

func TimeElapsed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before the request

		next.ServeHTTP(w, r)

		// After the request
	})
}
