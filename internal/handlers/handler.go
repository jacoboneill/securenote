package handlers

import "github.com/jacoboneill/securenote/internal/db"

type Handler struct {
	queries *db.Queries
}

func New(queries *db.Queries) *Handler {
	return &Handler{queries: queries}
}
