package main

import (
	"hostlii/apis"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signup", apis.SignUpHandler)
	http.HandleFunc("/login", apis.LoginHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
