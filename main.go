package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/joho/godotenv"

	"shitty-portfolio/internal/routes"
)

func main() {
	os.Setenv("APP_PORT", "3000")
	flagProd := flag.Bool("prod", false, "production state, default: false")
	flag.Parse()
	if *flagProd {
		os.Setenv("APP_PROD", "true")
	} else {
		os.Setenv("APP_PROD", "false")
	}

	// if err := godotenv.Load(envFilename); err != nil {
	// 	slog.Error("environment loading error:", "err", err)
	// }

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
		slog.Info("Server started at", "listenAddr", "localhost"+server.Addr, "prod", *flagProd)
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
