package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", HttpHandler)
	mux.HandleFunc("/flush", FlushHandler)
	handler := Logging(mux)

	Store()

	http.ListenAndServe(":"+port, handler)
}
