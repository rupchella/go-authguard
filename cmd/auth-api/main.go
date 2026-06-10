package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/katasuner/authguard/internal/config"
	apphttp "github.com/katasuner/authguard/internal/http"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config.Load()
	router := apphttp.NewRouter()

	logger.Info("starting auth api", "addr", cfg.HTTPAddr)

	if err := http.ListenAndServe(cfg.HTTPAddr, router); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
