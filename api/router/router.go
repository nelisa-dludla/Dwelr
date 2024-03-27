package router

import (
	"dwelr/api/handlers"

	"github.com/go-chi/chi/v5"
)

func Api() chi.Router {

	router := chi.NewRouter()
	// Listing routes
	router.Get("/listings", handlers.GetListings)
	router.Post("/listings", handlers.AddListing)
	router.Get("/listings/{id}", handlers.GetListing)
	router.Put("/listings/{id}", handlers.UpdateListing)
	router.Delete("/listings{id}", handlers.DeleteListing)
	// Search route
	router.Post("/search", handlers.SearchLocation)
	router.Post("/search/query", handlers.RetrieveCitySuburb)
	// User routes
	router.Get("/users/{id}", handlers.GetUser)
	router.Get("/users/{id}/listings", handlers.GetUserListings)
	// Location routes
	router.Post("/location", handlers.RetrieveLocation)

	return router
}
