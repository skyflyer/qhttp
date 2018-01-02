package main

import (
	"log"
	"net/http"
)

func logRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request [%s] %s\n", r.Method, r.URL)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
