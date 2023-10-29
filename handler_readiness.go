package main

import "net/http"

// this is how we define a http handler in go
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}