package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
