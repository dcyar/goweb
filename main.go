package main

import (
	"log"
	"net/http"

	c "github.com/dcyar/goweb/controllers"
	m "github.com/dcyar/goweb/middleware"
)

const PORT string = ":3000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", m.Log(c.Home))
	mux.HandleFunc("/contact", m.Log(c.Contact))

	log.Print("Running on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
