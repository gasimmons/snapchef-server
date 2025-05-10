package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./snapchef.db")
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	createTables()
}

func createTables() {
	query := `
    CREATE TABLE IF NOT EXISTS recipes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        ingredients TEXT
    );`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("❌ Failed to create table:", err)
	}
}
