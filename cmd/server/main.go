package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
	// "github.com/jacoboneill/securenote/internal/db"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if _, err := fmt.Fprintf(w, "Hello, world! %s", id); err != nil {
		log.Println(err)
	}
}

func main() {
	// Get Database connection
	conn, err := sql.Open("sqlite", "securenote.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("error closing db: %q", err)
		}
	}()

	// Run Migrations
	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("migrations applied")

	// Setup Handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// Start Server
	const addr = ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, logging(mux)))
}
