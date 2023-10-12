package main

import (
	"html/template"
	"net/http"
	"sync"
)

type UserData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

var (
	mu       sync.Mutex
	userData []UserData
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

			mu.Lock()
			userData = append(userData, UserData{Name: name, Email: email, Message: message})
			mu.Unlock()

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tmpl.Execute(w, userData)
	})

	http.ListenAndServe(":8080", nil)
}
