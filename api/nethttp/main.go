package main

import (
	"log/slog"
)

func main() {
	server := NewServer(newPetRepository())
	err := server.Run()

	if err != nil {
		slog.Error("fail to run server", slog.Any("err", err))
	}
}
