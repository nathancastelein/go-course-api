package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSONResult(w http.ResponseWriter, statusCode int, result interface{}) {
	b, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}
