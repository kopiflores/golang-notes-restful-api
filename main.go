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
