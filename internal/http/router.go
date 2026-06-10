package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/katasuner/authguard/internal/health"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Handler)

	return r
}
