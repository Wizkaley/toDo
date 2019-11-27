package routes

import (
	"todo/handlers"

	"github.com/go-chi/chi"
)

// Routes holds all the routes this app exposes as API's
func Routes() *chi.Mux {
	r := chi.NewRouter()
	// r.Route()
	//r.Use(middleware.RealIP())
	r.Get("/api/v1/list", handlers.GetArticles())
	return r
}
