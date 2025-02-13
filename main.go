<<<<<<< HEAD
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
=======
package main

import (
	"fmt"
	"golang-rest-api-notes/models"
	"golang-rest-api-notes/routes"
	"log"
	"net/http"
)

func main() {
	models.InitializeDB()

	router := routes.SetupRouter

	fmt.Println("Server starting on port : 8000")
	log.Fatal(http.ListenAndServe(":8000", router()))
}
>>>>>>> c6ab4a0 (update note handler)
