package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gasimmons/snapchef-server/api"
	"github.com/gasimmons/snapchef-server/db"
)

func main() {
	db.InitDB()

	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/recipes", api.RecipesHandler)
	http.HandleFunc("/users", api.UsersHandler)

	fmt.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
