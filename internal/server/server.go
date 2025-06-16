package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi"
)

type Server struct {
	Log        *log.Logger
	HttpServer *http.Server
}

func NewServer(log *log.Logger) *Server {
	r := chi.NewRouter()
	r.Get("/", handlers.IndexHandler)
	r.Post("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8800",
		Handler:      r,
		ErrorLog:     log,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	return &Server{Log: log, HttpServer: httpServer}
}
