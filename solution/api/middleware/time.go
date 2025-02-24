package main

import (
	"log/slog"
	"net/http"
	"time"
)

func TimeElapsed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()

		next.ServeHTTP(w, r)

		slog.Info("elapsed", slog.Any("time", time.Since(startedAt)))
	})
}
