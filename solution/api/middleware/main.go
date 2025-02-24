package main

import (
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!\n"))
}

func main() {
	http.Handle("GET /hello", TimeElapsed(Logger(http.HandlerFunc(HelloWorld))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
