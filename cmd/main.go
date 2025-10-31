package main

import "net/http"

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	server := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: handler,
	}

	server.ListenAndServe()
}
