package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, world!"); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func main() {
	const addr = ":8080"
	log.Printf("Listening on %s", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, logging(mux)))
}
