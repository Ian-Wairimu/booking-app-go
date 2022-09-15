package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

// main is the entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("started server at port %s", port))

	_ = http.ListenAndServe(port, nil)
}
