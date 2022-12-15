package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Name    string
	Email   string
	Message string
}

func Home(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/index.html"))

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/contact.html"))

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)

		return
	}

	details := ContactDetails{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	log.Printf("%+v", details)

	http.Redirect(w, r, "/contact", http.StatusSeeOther)
	tmpl.Execute(w, struct{ Success bool }{true})
}
