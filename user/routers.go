package user

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func RoutersHandler(memory *Memory) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				next.ServeHTTP(w, r)
			})
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			users, err := memory.FindAll()
			if err != nil {
				slog.Error("Error getting all users", "error", err)
				sendJSONErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			sendJSONSuccessResponse(w, http.StatusOK, users)
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var input MemoryUserInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				slog.Error("Error decoding request body", "error", err)
				sendJSONErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			if strings.TrimSpace(input.FirstName) == "" {
				slog.Error("First name is required")
				sendJSONErrorResponse(w, http.StatusBadRequest, "first_name is required")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if strings.TrimSpace(input.LastName) == "" {
				slog.Error("Last name is required")
				sendJSONErrorResponse(w, http.StatusBadRequest, "last_name is required")
				return
			}
			if strings.TrimSpace(input.Biography) == "" {
				slog.Error("Biography is required")
				sendJSONErrorResponse(w, http.StatusBadRequest, "biography is required")
				return
			}
			u, err := memory.Insert(input); if err != nil {
				slog.Error("Error inserting user", "error", err)
				sendJSONErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			sendJSONSuccessResponse(w, http.StatusCreated, u)
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			u, err := memory.FindById(id)
			if err != nil {
				slog.Error(fmt.Sprintf("Error getting user with id %s", id), "error", err)
				sendJSONErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			sendJSONSuccessResponse(w, http.StatusOK, u)
		})

		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			var input MemoryUserInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				slog.Error("Error decoding request body", "error", err)
				sendJSONErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			if strings.TrimSpace(input.FirstName) == "" {
				slog.Error("First name is required")
				sendJSONErrorResponse(w, http.StatusBadRequest, "first_name is required")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if strings.TrimSpace(input.LastName) == "" {
				slog.Error("Last name is required")
				sendJSONErrorResponse(w, http.StatusBadRequest, "last_name is required")
				return
			}
			if strings.TrimSpace(input.Biography) == "" {
				slog.Error("Biography is required")
				sendJSONErrorResponse(w, http.StatusBadRequest, "biography is required")
				return
			}
			id := chi.URLParam(r, "id")
			u := User{
				ID:        id,
				FirstName: input.FirstName,
				LastName:  input.LastName,
				Biography: input.Biography,
			}
			updatedUser, err := memory.Update(id, u)
			if err != nil {
				slog.Error(fmt.Sprintf("Error updating user with id %s", id), "error", err)
				sendJSONErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			sendJSONSuccessResponse(w, http.StatusOK, updatedUser)
		})

		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			u, err := memory.Delete(id)
			if err != nil {
				slog.Error(fmt.Sprintf("Error deleting user with id %s", id), "error", err)
				sendJSONErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			sendJSONSuccessResponse(w, http.StatusOK, u)
		})
	}
}