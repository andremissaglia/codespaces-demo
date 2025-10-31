package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	server := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: handler,
	}
	fmt.Println("Starting server on :8000")
	fmt.Println("Open http://localhost:8000")

	server.ListenAndServe()
}
