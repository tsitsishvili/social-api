package main

import (
	"log"

	"github.com/tsitsishvili/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
