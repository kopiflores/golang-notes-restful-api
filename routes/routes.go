package routes

import (
	"golang-rest-api-notes/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	router.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	router.HandleFunc("/notes/{id}", handlers.GetNoteByID).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	router.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

	return router
}
