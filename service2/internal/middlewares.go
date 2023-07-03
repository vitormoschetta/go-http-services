package internal

import (
	"log"
	"net/http"
)

func AcceptJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AcceptJSON")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
