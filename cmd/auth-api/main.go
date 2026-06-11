package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/rupchella/authguard/internal/config"
	appdb "github.com/rupchella/authguard/internal/db"
	apphttp "github.com/rupchella/authguard/internal/http"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config.Load()
	ctx := context.Background()
	dbPool, err := appdb.NewPool(ctx, cfg.DatabaseURL)

	if err != nil {
		logger.Error("failer to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	router := apphttp.NewRouter(dbPool)
	logger.Info("starting auth api", "addr", cfg.HTTPAddr)

	if err := http.ListenAndServe(cfg.HTTPAddr, router); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
