package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yunuskilicdev/yemeksepeti/service"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		GetKeyHandler(w, r)
	case "POST":
		PutKeyHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryKeyParam := queryParams["key"]
	k := queryKeyParam[0]
	store := service.Store()
	v := store.Get(k)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func PutKeyHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryKeyParam := queryParams["key"]
	k := queryKeyParam[0]
	queryValueParam := queryParams["value"]
	v := queryValueParam[0]
	store := service.Store()
	store.Put(k, v)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func FlushHandler(w http.ResponseWriter, r *http.Request) {
	store := service.Store()
	store.DeleteAll()
}
