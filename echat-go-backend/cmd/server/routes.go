package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// HandlerContext contains the exact dependencies your HTTP layer needs.
type HandlerContext struct {
	Logger *log.Logger
	// DB   *sql.DB
}

// New sets up the context for your handlers.
func New(logger *log.Logger) *HandlerContext {
	return &HandlerContext{
		Logger: logger,
	}
}

// Routes builds the chi router and registers paths.
func (h *HandlerContext) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/status", h.HandleStatus)

	return r
}

// HandleStatus is a method on our context, giving it clean access to h.Logger.
func (h *HandlerContext) HandleStatus(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println("Status endpoint reached")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
