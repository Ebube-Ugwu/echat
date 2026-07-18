package server

import (
	"errors"
	"log"
	"net/http"
	"time"
)

type Server struct {
	addr   string
	logger *log.Logger // The server stores a pointer to the shared logger
	router http.Handler
}

// New accepts the logger from the application setup
func New(addr string, logger *log.Logger, router http.Handler) *Server {
	return &Server{
		addr:   addr,
		logger: logger,
		router: router,
	}
}

func (s *Server) Serve() error {
	srv := &http.Server{
		Addr:         s.addr,
		Handler:      s.router,
		ErrorLog:     s.logger, // Reuses the app logger for internal HTTP server errors
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Uses the shared logger for operational logs
	s.logger.Printf("Starting server on %s", s.addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
