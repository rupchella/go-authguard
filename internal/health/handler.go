package health

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Response struct {
	Status   string `json:"status"`
	Database string `json:"database,omitempty"`
}

type Handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Response{
		Status: "ok",
	})
}

func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	if err := h.db.Ping(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, Response{
			Status:   "not_ready",
			Database: "down",
		})
		return
	}

	writeJSON(w, http.StatusOK, Response{
		Status:   "ready",
		Database: "ok",
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(response)

}
