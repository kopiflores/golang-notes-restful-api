<<<<<<< HEAD
// models/note.go
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" 
)

// Note struct defines the note entity
type Note struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Initialize DB connection (optional)
var DB *gorm.DB

func InitializeDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "./notes.db")
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Note{})
}
=======
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Agenda struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Time   string `json:"time"`
	Agenda string `json:"agenda"`
}

var DB *gorm.DB

func InitializeDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "./notes.db")
	DB.LogMode(true)
	if err != nil {
		panic("Failed to connect database")
	}

	DB.AutoMigrate(&Note{})
	DB.AutoMigrate(&Agenda{})
}
>>>>>>> c6ab4a0 (update note handler)
