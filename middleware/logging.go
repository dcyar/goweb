package middleware

import (
	"log"
	"net/http"
)

func Log(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.URL.Path, r.Method)
		f(w, r)
	}
}
