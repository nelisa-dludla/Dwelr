package web_routes

import (
	"dwelr/web/handlers"

	"github.com/go-chi/chi/v5"
)

func PagesRoutes(r chi.Router) {
	r.Get("/", handlers.RenderHome)
	r.Get("/about", handlers.RenderAbout)
	r.Get("/login", handlers.RenderLogin)
	r.Get("/register", handlers.RenderRegister)
	r.Get("/results", handlers.RenderResult)
	r.Get("/view-listing", handlers.RenderViewListing)
	r.Get("/profile", handlers.RenderProfile)
	r.Get("/add-listing", handlers.RenderAddListing)
}
