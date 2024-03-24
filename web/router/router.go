package router

import (
	apiRouter "dwelr/api/router"
	"dwelr/controllers"
	"dwelr/web/handlers"
	auth "dwelr/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	key = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func New() chi.Router {
	router := chi.NewRouter()
	// Middleware
	router.Use(middleware.Logger)
	router.Use(auth.RequireAuth)
	// Authentication middlware
	// Serve static files
	router.Handle("/views/css/*", http.StripPrefix("/views/css", http.FileServer(http.Dir("./views/css"))))
	router.Handle("/views/images/*", http.StripPrefix("/views/images", http.FileServer(http.Dir("./views/images"))))
	router.Handle("/views/scripts/*", http.StripPrefix("/views/scripts", http.FileServer(http.Dir("./views/scripts"))))
	// APIs
	router.Mount("/api", apiRouter.Api())
	// Serve pages
	router.Get("/", handlers.Index)
	router.Get("/about", handlers.About)
	router.Get("/login", handlers.Login)
	router.Post("/login", controllers.Login)
	router.Get("/register", handlers.RegisterPage)
	router.Post("/register", controllers.Register)
	router.Get("/home", handlers.Home)
	router.Get("/view-listing", handlers.ViewListing)
	router.Get("/add-listing", handlers.AddListing)
	router.Get("/validate", controllers.Validate)
	router.NotFound(handlers.Not_Found)
	// Auth

	return router
}
