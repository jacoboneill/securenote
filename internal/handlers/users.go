package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jacoboneill/securenote/internal/db"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := h.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        "jacoboneill2000@outlook.com",
		PasswordHash: "hash123",
	})
	if err != nil {
		log.Printf("error: %q", err)
	}

	if _, err := fmt.Fprint(w, id); err != nil {
		log.Println(err)
	}
}
