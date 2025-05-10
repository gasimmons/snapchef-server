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

	http.HandleFunc("/recipes", api.RecipesHandler)

	fmt.Println("🚀 Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
