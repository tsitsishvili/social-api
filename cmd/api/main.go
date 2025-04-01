package main

import (
	"github.com/joho/godotenv"
	"github.com/tsitsishvili/social/internal/db"
	"github.com/tsitsishvili/social/internal/env"
	"github.com/tsitsishvili/social/internal/store"
	"go.uber.org/zap"
	"time"
)

const version = "1.0.0"

//	@title			Social API
//	@description	API for Social app
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	godotenv.Load()

	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", "postgres://postgres:root@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 5),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 2),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env:    env.GetString("ENV", "local"),
		apiURL: env.GetString("API_URL", "localhost:4000"),
		mail: mailConfig{
			exp: env.GetDuration("MAIL_EXP", 24*time.Hour),
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.dsn,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("Database connection established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
