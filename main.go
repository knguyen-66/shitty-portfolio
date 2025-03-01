package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/joho/godotenv"

	"shitty-portfolio/data"
	"shitty-portfolio/internal/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("environment loading error:", "err", err)
	}

	err := data.InitDatabase(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		slog.Error("database initialization error:", "err", err)
		os.Exit(1)
	}
	defer data.CloseDatabase()

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.LoadRouterPaths(r)

	server := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: r,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("Server started at", "listenAddress", "localhost"+server.Addr, "prod", os.Getenv("APP_PRODUCTION"))
		err := server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("Server shutdown complete")
		} else if err != nil {
			slog.Error("Server error, exiting;", "err", err)
			os.Exit(1)
		}
	}()

	// wait for termination signal
	<-sig

	slog.Info("Server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "err", err.Error())
		os.Exit(1)
	}
}
