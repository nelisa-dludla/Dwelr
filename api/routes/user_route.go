package routes

import (
	"dwelr/api/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Get("/users/{id}", handlers.GetUser)
	r.Get("/users/{id}/listings", handlers.GetUserListings)
}
