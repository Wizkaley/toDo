package main

import (
	"log"
	"net/http"
	"todo/routes"
)

func main() {
	//r := chi.NewRouter()
	r := routes.Routes()
	log.Fatal(http.ListenAndServe(":3000", r))
}
