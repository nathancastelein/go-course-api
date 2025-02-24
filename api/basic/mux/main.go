package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello/{name}", http.HandlerFunc(HelloWorld))

	http.ListenAndServe(":8080", mux)
}
