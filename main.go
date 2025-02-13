package main

import (
	"fmt"
	"golang-rest-api-notes/models"
	"golang-rest-api-notes/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func SetupCORS(handler http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Ganti "*" dengan origin spesifik jika diperlukan
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(handler)
}

func main() {
	models.InitializeDB()

	router := routes.SetupRouter()

	// Bungkus router dengan SetupCORS untuk menangani CORS
	handler := SetupCORS(router)

	fmt.Println("Server starting on port : 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
