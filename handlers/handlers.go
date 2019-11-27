package handlers

import (
	"github.com/go-chi/render"
	"net/http"
)

// GetArticles returns a list of articles
func GetArticles() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "Hellow From GetHandler")
	})
}
