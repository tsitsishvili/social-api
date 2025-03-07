package main

import (
	"log"

	"github.com/tsitsishvili/social/internal/db"
	"github.com/tsitsishvili/social/internal/env"
	"github.com/tsitsishvili/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", "postgres://postgres:root@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 5),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 2),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.dsn,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
