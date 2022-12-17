package middleware

import (
	"github.com/dcyar/goweb/routes"
	"html/template"
	"log"
	"net/http"
)

func Log(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.URL.Path, r.Method)

		rt := routes.Routes()
		_, exists := rt[r.URL.Path]

		if !exists {
			notFound(w, r)
			return
		}

		f(w, r)
	}
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/404.html"))

	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, nil)
}
