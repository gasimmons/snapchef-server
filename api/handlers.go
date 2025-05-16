package api

import (
	"encoding/json"
	"github.com/gasimmons/snapchef-server/db"
	"net/http"
)

func RecipesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		recipes, err := db.GetAllRecipes()
		if err != nil {
			http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipes)
	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		ingredients := r.FormValue("ingredients")
		if err := db.InsertRecipe(title, ingredients); err != nil {
			http.Error(w, "Failed to insert recipe", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
