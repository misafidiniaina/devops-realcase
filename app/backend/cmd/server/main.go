package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misafidiniaina/cloud-platform-lab/internal/server"
)

func main() {
	ctx := context.Background()
	databaseURL := os.Getenv("DATABASE_URL")
	var pool *pgxpool.Pool
	if databaseURL != "" {
		var err error
		pool, err = pgxpool.New(ctx, databaseURL)
		if err != nil {
			log.Fatal(err)
		}
		defer pool.Close()
	}

	mux := http.NewServeMux()
	server.New(pool).Register(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
