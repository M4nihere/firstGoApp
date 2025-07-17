package apis

import (
	"encoding/json"
	"os"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const dbPath = "data/users.json"

func LoadUsers() ([]User, error) {
	file, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(file, &users)
	return users, err
}

func SaveUsers(users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dbPath, data, 0644)
}
