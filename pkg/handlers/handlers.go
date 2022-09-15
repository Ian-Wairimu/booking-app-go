package handlers

import (
	"net/http"
	"wairimuian.com/booking-app-go/pkg/render"
)

//Home this is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RendersTemplate(w, "home.html")
}

// About this is the about page of the application
func About(w http.ResponseWriter, r *http.Request) {
	render.RendersTemplate(w, "about.html")
}
