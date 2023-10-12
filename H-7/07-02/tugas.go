package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	PhotoURL string `json:"website"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer response.Body.Close()

	var users []User
	err = json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		fmt.Println("Error while decoding JSON:", err)
		return
	}

	fmt.Printf("%-5s | %-30s | %-15s | %-30s | %-50s\n", "ID", "Email", "Phone", "Fullname", "Photo URL")
	fmt.Println(strings.Repeat("-", 110))
	for _, user := range users {
		fmt.Printf("%-5d | %-30s | %-15s | %-30s | %-50s\n", user.ID, user.Email, user.Phone, user.Name, user.PhotoURL)
	}
}
