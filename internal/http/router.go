package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rupchella/authguard/internal/health"
)

func NewRouter(db *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	healthHandler := health.NewHandler(db)

	r.Get("/health", healthHandler.Health)
	r.Get("/ready", healthHandler.Ready)

	return r
}
