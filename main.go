package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "welcome from chi")
	})

	http.ListenAndServe(":3000", r)
}
