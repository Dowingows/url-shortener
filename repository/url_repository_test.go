package repository

import (
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
)

func setupDB(t *testing.T) *sql.DB {
	connStr := "host=" + "localhost" +
		" port=" + "5432" +
		" user=" +  "postgres" +
		" password=" + "postgres" +
		" dbname=" + "shortener" +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Exec("DELETE FROM urls")
		db.Close()
	})

	return db
}

func TestURLRepository_Create(t *testing.T) {
	db := setupDB(t)
	repo := NewURLRepository(db)

	id, short, err := repo.Create("https://google.com")
	if err != nil {
		t.Fatalf("erro ao criar url: %v", err)
	}

	if id == 0 {
		t.Fatal("id não pode ser zero")
	}

	if short == "" {
		t.Fatal("short_url não pode ser vazio")
	}

	// Confirma no banco
	var original string
	err = db.QueryRow(
		"SELECT original_url FROM urls WHERE id = $1",
		id,
	).Scan(&original)

	if err != nil {
		t.Fatalf("erro ao buscar url no banco: %v", err)
	}

	if original != "https://google.com" {
		t.Fatalf("esperado %s, recebido %s", "https://google.com", original)
	}
}
