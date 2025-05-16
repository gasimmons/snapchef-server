package api

import (
	"encoding/json"
	"net/http"

	"github.com/gasimmons/snapchef-server/db"
)

type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser NewUser
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, err := db.CreateUser(
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		newUser.Password,
	)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"userId":    userID,
		"email":     newUser.Email,
		"firstName": newUser.FirstName,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
