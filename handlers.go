package main

import (
	"net/http"
)

//Home this is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About this is the about page of the application
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
