package main

import (
	"fmt"
	"net/http"
)

func setUpRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Starting server...")
	})
}

func main() {
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}
