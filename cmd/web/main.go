package main

import (
	"fmt"
	"net/http"
	"wairimuian.com/booking-app-go/pkg/handlers"
)

const port = ":8080"

// main is the entry point of the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("started server at port %s", port))

	_ = http.ListenAndServe(port, nil)
}
