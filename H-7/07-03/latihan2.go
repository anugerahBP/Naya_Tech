package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Contact struct {
	Name     string
	Location string
	Email    string
}

var contacts []Contact

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("template/home2.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.FormValue("name")
		location := r.FormValue("location")
		email := r.FormValue("email")

		contacts = append(contacts, Contact{Name: name, Location: location, Email: email})

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("template/contact2.html"))
		tmpl.Execute(w, contacts)
	}
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
