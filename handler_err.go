package main

import "net/http"

// this is how we define a http handler in go
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}