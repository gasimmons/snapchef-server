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
	userQuery := `CREATE TABLE IF NOT EXISTS users (
    	userId INTEGER PRIMARY KEY AUTOINCREMENT,
    	firstName TEXT,
    	lastName TEXT,
    	username TEXT,
    	email TEXT,
    	password_salt TEXT,
    	password_hash TEXT
	);`

	recipeQuery := `
    CREATE TABLE IF NOT EXISTS recipes (
        recipeId INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        ingredients TEXT
    );`

	userRecipeQuery := `CREATE TABLE IF NOT EXISTS user_recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    recipeId INTEGER NOT NULL,   
    FOREIGN KEY(userId) REFERENCES users(userId),
    FOREIGN KEY(recipeId) REFERENCES recipes(recipeId)            
	);`

	_, err := DB.Exec(userQuery)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	_, err = DB.Exec(recipeQuery)
	if err != nil {
		log.Fatal("Failed to create recipe table:", err)
	}

	_, err = DB.Exec(userRecipeQuery)
	if err != nil {
		log.Fatal("Failed to create user_recipes table:", err)
	}
}
