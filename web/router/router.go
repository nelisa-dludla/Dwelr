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

func New() chi.Router {
	router := chi.NewRouter()
	// Middleware
	router.Use(middleware.Logger)
	router.Use(auth.AuthStatus)
	// Serve static files
	router.Handle("/views/css/*", http.StripPrefix("/views/css", http.FileServer(http.Dir("./views/css"))))
	router.Handle("/views/images/*", http.StripPrefix("/views/images", http.FileServer(http.Dir("./views/images"))))
	router.Handle("/views/scripts/*", http.StripPrefix("/views/scripts", http.FileServer(http.Dir("./views/scripts"))))
	// Routes that don't require authentication
	router.Group(func(r chi.Router) {
		r.Get("/", handlers.Index)
		r.Get("/login", handlers.Login)
		r.Post("/login", controllers.Login)
		r.Get("/register", handlers.RegisterPage)
		r.Post("/register", controllers.Register)
		r.NotFound(handlers.Not_Found)
	})
	// that require an authentication check
	router.Group(func(r chi.Router) {
		r.Get("/about", handlers.About)
		r.Get("/view-listing/{id}", handlers.ViewListing)
		r.Get("/add-listing", handlers.AddListing)
		r.Get("/profile/{id}", handlers.RenderProfile)
		r.Get("/logout", controllers.Logout)
	})
	// APIs
	router.Mount("/api", apiRouter.Api())
	//router.Get("/validate", controllers.Validate)

	return router
}
