package main

import (
	"net/http"
	"os"

	"github.com/yunuskilicdev/yemeksepeti/service"
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

	service.Store()

	http.ListenAndServe(":"+port, handler)
}
