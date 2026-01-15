package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type Payload struct {
		URL string `json:"url"`
	}

	var pay Payload
	if err := json.NewDecoder(r.Body).Decode(&pay); err != nil || pay.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type Response struct{
		Short string `json:"short_url"`
		Original string `json:"original_url"`
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	
	json.NewEncoder(w).Encode(Response{
		Short: "asd",
		Original: pay.URL,
	})
	
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" || code == "shorten" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url := "https://google.com"

	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.HandleFunc("/hello", pingHandler)
	http.HandleFunc("/shortener", shortenHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}