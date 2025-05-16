package api

import (
	"encoding/json"
	"github.com/gasimmons/snapchef-server/db"
	"net/http"
)

type newRecipe struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	UserId      int    `json:"userId"`
}

func RecipesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID := r.URL.Query().Get("userId")
		if userID == "" {
			http.Error(w, "Missing userId", http.StatusBadRequest)
			return
		}

		recipes, err := db.GetRecipesByUser(userID)
		if err != nil {
			http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(recipes)
		if err != nil {
			return
		}

	case http.MethodPost:
		var NewRecipe newRecipe
		if err := json.NewDecoder(r.Body).Decode(&NewRecipe); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err := db.InsertRecipeWithUser(NewRecipe.Title, NewRecipe.Ingredients, NewRecipe.UserId)
		if err != nil {
			http.Error(w, "Failed to insert recipe", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
