package main

import (
	"log/slog"
	"net/http"
)

/*
Using json.Marshal, encode the result in parameter

If you encounter an error, write an status code 500, with the text "Internal Server Error".

If you have no error, then:
- write the status code 200
- write a header "Content-Type" with the value "application/json"
- write the encoded result
*/

func WriteJSONResult(w http.ResponseWriter, statusCode int, result interface{}) {
	slog.Debug("Nothing to do...")
}
