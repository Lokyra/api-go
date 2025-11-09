package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	fmt.Println("Server listening to :5000")

	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
