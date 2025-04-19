package user

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/errors"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/server"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

// NewHandler は新しいユーザーハンドラーを生成します。
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", server.WithErrorHandler(h.ListUsers))
		r.Post("/", server.WithErrorHandler(h.CreateUser))
		r.Get("/{id}", server.WithErrorHandler(h.GetUser))
		r.Put("/{id}", server.WithErrorHandler(h.UpdateUser))
		r.Delete("/{id}", server.WithErrorHandler(h.DeleteUser))
	})
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.NewBadRequestError("invalid_id", "Invalid user ID format")
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewNotFoundError("user_not_found", "User not found")
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var input CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return errors.NewBadRequestError("invalid_input", "Could not parse request body")
	}

	if input.Name == "" {
		return errors.NewBadRequestError("invalid_input", "Name is required")
	}

	user, err := h.service.CreateUser(r.Context(), input)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.NewBadRequestError("invalid_id", "Invalid user ID format")
	}

	var input UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return errors.NewBadRequestError("invalid_input", "Could not parse request body")
	}

	user, err := h.service.UpdateUser(r.Context(), id, input)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewNotFoundError("user_not_found", "User not found")
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.NewBadRequestError("invalid_id", "Invalid user ID format")
	}

	err = h.service.DeleteUser(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
