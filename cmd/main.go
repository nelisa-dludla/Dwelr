package main

import (
	"dwelr/api/routes"
	"dwelr/db"
	"dwelr/web/web_routes"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Initilize database
	dsn := fmt.Sprintf("postgres://%s:%s@localhost/dwelr_db", dbUser, dbPass)
	_, err := db.InitDB(dsn)
	if err != nil {
		fmt.Println("Error initializing Database:", err)
	}
	// Handle routes
	router := chi.NewRouter()
	router.Use(middleware.Logger)	
	// Serve CSS files
	router.Handle("/views/css/*", http.StripPrefix("/views/css", http.FileServer(http.Dir("./views/css"))))
	router.Handle("/views/images/*", http.StripPrefix("/views/images", http.FileServer(http.Dir("./views/images"))))
	routes.UserRoutes(router)
	routes.ListingRoutes(router)
	web_routes.PagesRoutes(router)

	fmt.Println("Starting server on port 3000")
	serverErr := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Failed to start the server:", serverErr)
	}
}
