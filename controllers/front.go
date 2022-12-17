package controllers

import (
	"context"
	"github.com/dcyar/goweb/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"log"
	"net/http"
	"time"
)

type ContactDetails struct {
	Name      string
	Email     string
	Message   string
	CreatedAt time.Time
}

func Home(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/index.html"))

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/contact.html"))
	client := config.DatabaseConnect()
	defer config.DatabaseClose(client)
	coll := client.Database(config.Database).Collection("contacts")

	if r.Method != http.MethodPost {
		cursor, err := coll.Find(context.TODO(), bson.D{{}}, options.Find().SetSort(bson.D{{"createdat", -1}}))
		if err != nil {
			panic(err)
		}

		var contacts []ContactDetails
		if err = cursor.All(context.TODO(), &contacts); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, struct{ Contacts []ContactDetails }{contacts})

		return
	}

	details := ContactDetails{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		Message:   r.FormValue("message"),
		CreatedAt: time.Now(),
	}

	newContact, err := coll.InsertOne(context.TODO(), details)

	if err != nil {
		panic(err)
	}

	log.Printf("%+v", newContact)

	http.Redirect(w, r, "/contact", http.StatusSeeOther)
	tmpl.Execute(w, struct{ Success bool }{true})
}
