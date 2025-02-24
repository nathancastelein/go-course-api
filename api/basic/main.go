package main

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!\n"))
}

func main() {
	http.HandleFunc("GET /hello", HelloWorld)

	http.ListenAndServe(":8080", nil)
}
