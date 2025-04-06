package main

import (
	"github.com/joho/godotenv"
	"github.com/tsitsishvili/social/internal/db"
	"github.com/tsitsishvili/social/internal/env"
	"github.com/tsitsishvili/social/internal/store"
	"log"
)

func main() {
	godotenv.Load()

	addr := env.GetString("DB_DSN", "postgres://postgres:root@localhost:5432/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer conn.Close()

	storage := store.NewStorage(conn)

	db.Seed(storage, conn)
}
