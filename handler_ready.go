package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
