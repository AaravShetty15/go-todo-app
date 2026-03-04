package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Todo API server...")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	http.ListenAndServe(":8080", nil)
}