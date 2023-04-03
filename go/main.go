package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	http.HandleFunc("/ping", handler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	hash, _ := hashPassword("test")
	w.Write([]byte("A hashed password for 'test' " + hash))
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
