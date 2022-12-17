package main

import (
	c "github.com/dcyar/goweb/controllers"
	m "github.com/dcyar/goweb/middleware"
	"github.com/dcyar/goweb/routes"
	ws "github.com/dcyar/goweb/websocket"
	"log"
	"net/http"
)

const PORT string = ":3000"

func main() {
	me := ws.New()
	defer me.Close()

	mux := http.NewServeMux()
	r := routes.Routes()

	mux.HandleFunc(r["/"], m.Log(c.Home))
	mux.HandleFunc(r["/contact"], m.Log(c.Contact))
	mux.HandleFunc(r["/ws"], m.Log(func(w http.ResponseWriter, r *http.Request) {
		me.HandleRequest(w, r)
	}))

	c.HandleContactMessage(me)

	log.Print("Running on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
