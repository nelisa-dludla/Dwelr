package routes

import (
	"dwelr/api/handlers"

	"github.com/go-chi/chi/v5"
)

func ListingRoutes(r chi.Router) {
	r.Get("/listings", handlers.GetListings)
	r.Post("/listings", handlers.AddListing)
	r.Get("/listings/{id}", handlers.GetListing)
	r.Put("/listings/{id}", handlers.UpdateListing)
	r.Delete("/listings{id}", handlers.DeleteListing)
}
