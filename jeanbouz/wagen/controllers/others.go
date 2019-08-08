package controllers

// This file handles the http queries and the display

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	// mainPath is the path of the wagen folder
	// It will be useful to access quickly all the files (eg. templates)
	mainPath string = "C:/Go/src/github.com/jeanbouz/wagen/templates/"
)

var (
	// Templates is used to parse all the templates once at program initialization (and avoid to call ParseFiles each time we need one)
	// Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil
	// A panic is appropriate here, if the templates cannot be loaded we need to exit the program
	// The function template.ParseFiles will read the contents of example.html and return a (*Template, error)
	Templates = template.Must(template.ParseFiles(mainPath+"edit.html",
		mainPath+"login.html",
		mainPath+"cars.html",
		mainPath+"customers.html",
		mainPath+"users.html"))
)

// RootHandler makes the web root redirect to the view connection
func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

// LoginHandler handles the view connection
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Gets request method
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// Execute Templates is the same than t := template.ParseFiles + t.Execute (because we alreadu parsed all the files !)
		templates.ExecuteTemplate(w, "login.html", nil)
	} else {
		r.ParseForm()
		// Logic part of login
		chain := "Welcome " + r.Form["username"][0] + ", your password is :" + r.Form["password"][0]
		fmt.Fprintf(w, chain)
	}
}

// CarsHandler handles the view cars
func CarsHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "cars.html", nil)
}

// CustomersHandler handles the view cars
func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "consumers.html", nil)
}

// UsersHandler handles the view cars
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "users.html", nil)
}
