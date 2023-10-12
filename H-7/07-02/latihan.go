package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	PhotoURL string `json:"website"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	url := "https://jsonplaceholder.typicode.com/users"

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error while making the request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var users []User
	err = json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Error while decoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)

	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
