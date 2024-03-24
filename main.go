package main

import (
	"dwelr/setup"
	"dwelr/web/router"
	"fmt"
	"net/http"
	"os"
)

func init() {
	setup.LoadEnv()
	db := setup.ConnectToDb()
	setup.SyncDb(db)
}

func main() {
	port := os.Getenv("PORT")
	// Router
	r := router.New()
	// Start server
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
