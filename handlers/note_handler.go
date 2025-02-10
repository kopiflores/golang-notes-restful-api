// handlers/note_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"golang-rest-api-notes/models"
	"github.com/gorilla/mux"
)

// Get all notes
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	models.DB.Find(&notes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// Create a new note
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	models.DB.Create(&note)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// Get a single note by ID
func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var note models.Note
	if err := models.DB.First(&note, params["id"]).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// Update an existing note
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var note models.Note
	if err := models.DB.First(&note, params["id"]).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&note)
	models.DB.Save(&note)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// Delete a note
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)	
	var note models.Note
	if err := models.DB.First(&note, params["id"]).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	models.DB.Delete(&note)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Note successfully deleted"})
}