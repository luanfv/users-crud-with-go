package main

import (
	"log/slog"
	"net/http"
	"time"

	"userCrud/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	memory := user.NewMemory()
	server := http.Server{
		Addr:    ":8080",
		Handler: handlers(memory),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
	}
	err := server.ListenAndServe(); if err != nil {
		slog.Error("Server cannot run", "error", err)
		return
	}
	slog.Info("Server is offline")
}

func handlers(memory *user.Memory) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api/users", user.RoutersHandler(memory))
	return r
}