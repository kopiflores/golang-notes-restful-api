// routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
	"golang-rest-api-notes/handlers"
)

// SetupRouter initializes routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id}", handlers.GetNoteByID).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")
	return r
}