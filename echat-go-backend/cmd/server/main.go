package main

import (
	"database/sql"
	"github.com/ebubeugwu/echat/echat-go-backend/internal/server"
	"log"
)

// In main.go (or internal/app)
type Application struct {
	Logger *log.Logger
	DB     *sql.DB
	// Config, Mailer, etc.
}

func main() {
	app := &Application{
		Logger: log.Default(),
	}

	// Pass the app or just its handlers to the server
	srv := server.New(app)
	srv.Serve()
}
