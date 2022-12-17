package main

import (
	"github.com/dcyar/goweb/routes"
	"log"
	"net/http"

	c "github.com/dcyar/goweb/controllers"
	m "github.com/dcyar/goweb/middleware"
)

const PORT string = ":3000"

func main() {
	mux := http.NewServeMux()
	r := routes.Routes()

	mux.HandleFunc(r["/"], m.Log(c.Home))
	mux.HandleFunc(r["/contact"], m.Log(c.Contact))

	log.Print("Running on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
