package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"shitty-portfolio/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	listenAddr := ":3000"

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./internal/static"))))
	r.Get("/", handlers.Make(handlers.HandleHello))

	server := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("Server started at", "listenAddr", " localhost"+listenAddr)
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "err", err.Error())
		os.Exit(1)
	}
}
