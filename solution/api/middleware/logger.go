package main

import (
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("starting request")

		next.ServeHTTP(w, r)

		slog.Info("request end")
	})
}
