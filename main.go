// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"golang-rest-api-notes/models"
	"golang-rest-api-notes/routes"
)

func main() {
	// Initialize DB
	models.InitializeDB()

	// Setup routes
	r := routes.SetupRouter()

	// Start server
	fmt.Println("Starting server on :8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}