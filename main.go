package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"userCrud/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type response struct {
	Data any `json:"data"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	memory := user.NewMemory()
	http := http.Server{
		Addr:    ":8080",
		Handler: handlers(memory),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
	}
	err := http.ListenAndServe(); if err != nil {
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

	r.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		users, err := memory.GetAll(); if err != nil {
			slog.Error("Error getting all users", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(response{Data: users})
	})

	r.Post("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var input user.MemoryInsertInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			slog.Error("Error decoding request body", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Print(input)
		id, err := memory.Insert(input); if err != nil {
			slog.Error("Error inserting user", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(response{Data: map[string]string{"id": id}})
	})

	return r
}