package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetDatabase() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./bookmarks.db")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		// Initialize the database
		initDatabase()
	}
	return db
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}

func initDatabase() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS bookmarks (
		id TEXT PRIMARY KEY, 
		url TEXT NOT NULL,
		created_at DATETIME
	);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create bookmarks table: %v", err)
	}
}
