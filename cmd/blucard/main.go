package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/blucard/blucard/internal/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize the server: %v", err)
	}

	go func() {
		if err := srv.HTTP.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Received SIGINT, starting teardown...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.HTTP.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown server: %v", err)
	}
	log.Println("Shutdown complete.")
}
