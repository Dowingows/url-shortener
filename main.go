package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Dowingows/url-shortener/repository"
	_ "github.com/lib/pq"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func shortenHandler(repo *repository.URLRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		_, shortCode, err := repo.Create(pay.URL)

		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		type Response struct {
			Short    string `json:"short_url"`
			Original string `json:"original_url"`
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(Response{
			Short:    "http://host.docker.internal:8080/" + shortCode,
			Original: pay.URL,
		})

	}
}

func redirectHandler(repo *repository.URLRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimPrefix(r.URL.Path, "/")
		if code == "" || code == "shorten" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		originalUrl, err := repo.Find(code)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Print(err)
			return
		}

		if originalUrl == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalUrl, http.StatusFound)
	}
}

func connectDB() *sql.DB {
	connStr := "host=postgres port=5432 user=postgres password=postgres dbname=shortener sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Erro ao abrir banco: %v", err))
	}

	if err := db.Ping(); err != nil {
		fmt.Print(fmt.Sprintf("Erro ao conectar no banco: %v", err))
	}

	return db
}

func main() {

	db := connectDB()
	repo := repository.NewURLRepository(db)

	http.HandleFunc("/hello", pingHandler)
	http.HandleFunc("/shortener", shortenHandler(repo))
	http.HandleFunc("/{code}", redirectHandler(repo))

	fmt.Println("Servidor rodando em http://host.docker.internal:8080")
	http.ListenAndServe(":8080", nil)
}
