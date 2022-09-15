package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

//Home this is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About this is the about page of the application
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
func renderTemplate(w http.ResponseWriter, html string) {
	parsedFile, _ := template.ParseFiles("./templates/" + html)
	err := parsedFile.Execute(w, nil)
	if err != nil {
		fmt.Println("failed to parse the html template")
		return
	}
}

// main is the entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("started server at port %s", port))

	_ = http.ListenAndServe(port, nil)
}
