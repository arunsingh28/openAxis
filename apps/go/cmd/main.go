package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arunsingh28/openaxis/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})

	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	slog.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	// channel to listen for server shutdown
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// routine to start the server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server: %f", err.Error())
		}
	}()

	<-done

	slog.Info("server is shutting down", slog.String("address", cfg.HTTPServer.Address))

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(context)
	if err != nil {
		slog.Info("failed to shutdown server", slog.String("address", cfg.HTTPServer.Address), slog.String("error", err.Error()))
	}

	slog.Info("server is shutdown", slog.String("address", cfg.HTTPServer.Address))
}
