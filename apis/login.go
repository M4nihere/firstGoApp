package apis

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	users, err := LoadUsers()
	if err != nil {
		http.Error(w, "Error reading users", http.StatusInternalServerError)
		return
	}

	for _, u := range users {
		if u.Email == creds.Email && u.Password == creds.Password {
			w.Write([]byte("Login successful"))
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
