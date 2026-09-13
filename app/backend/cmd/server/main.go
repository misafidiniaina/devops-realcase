package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/misafidiniaina/cloud-platform-lab/internal/config"
	"github.com/misafidiniaina/cloud-platform-lab/internal/database"
	api "github.com/misafidiniaina/cloud-platform-lab/internal/server"
)

func main() {
	cfg := config.Load()
	db, err := database.Connect(context.Background(), cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer db.Close()
	if err := database.Migrate(context.Background(), db); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	srv := &http.Server{Addr: ":" + cfg.AppPort, Handler: api.NewRouter(db, cfg), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		log.Printf("server listening on :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
