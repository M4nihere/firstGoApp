package apis

import (
	"encoding/json"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	users, err := LoadUsers()
	if err != nil {
		http.Error(w, "Error reading users", http.StatusInternalServerError)
		return
	}

	// check if user already exists
	for _, u := range users {
		if u.Email == newUser.Email {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
	}

	users = append(users, newUser)
	if err := SaveUsers(users); err != nil {
		http.Error(w, "Could not save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
