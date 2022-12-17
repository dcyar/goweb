package controllers

import (
	"context"
	"encoding/json"
	"github.com/dcyar/goweb/config"
	"github.com/olahol/melody"
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
}

func HandleContactMessage(me *melody.Melody) {
	me.HandleMessage(func(s *melody.Session, msg []byte) {
		client := config.DatabaseConnect()
		defer config.DatabaseClose(client)

		coll := client.Database(config.Database).Collection("contacts")

		tempCd := ContactDetails{}
		err := json.Unmarshal(msg, &tempCd)

		if err != nil {
			log.Println(err)
		}

		tempCd.CreatedAt = time.Now()

		_, err = coll.InsertOne(context.TODO(), tempCd)

		if err != nil {
			panic(err)
		}

		me.Broadcast(msg)
	})
}
