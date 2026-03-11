package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	db "github.com/jacoboneill/securenote/internal/db"
	"github.com/jacoboneill/securenote/internal/handlers"
	_ "modernc.org/sqlite"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func initDB(dataSourceName string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", dataSourceName)
	return conn, err
}

func runMigrations(conn *sql.DB, sourceURL string) error {
	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "sqlite", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func initMux(queries *db.Queries) http.Handler {
	h := handlers.New(queries)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", h.CreateUser)

	return logging(mux)
}

func main() {
	conn, err := initDB("securenote.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("error closing db: %q", err)
		}
	}()

	queries := db.New(conn)

	if err := runMigrations(conn, "file://migrations"); err != nil {
		log.Fatal(err)
	}

	mux := initMux(queries)

	// Start Server
	const addr = ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
