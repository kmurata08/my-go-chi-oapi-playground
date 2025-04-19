package user

import (
	"encoding/json"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/errors"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/server"
	user2 "github.com/kmurata08/my-go-chi-oapi-playground/internal/gen/user"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// ErrorAdapter ServerInterfaceに適合するアダプター
type ErrorAdapter struct {
	*Handler
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		return err
	}

	response := map[string]interface{}{
		"users": users,
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request, id int) error {
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
	var input user2.CreateUserRequest
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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request, id int) error {
	var input user2.UpdateUserRequest
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

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) error {
	err := h.service.DeleteUser(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (a *ErrorAdapter) ListUsers(w http.ResponseWriter, r *http.Request) {
	server.WithErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return a.Handler.ListUsers(w, r)
	})(w, r)
}

func (a *ErrorAdapter) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	server.WithErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return a.Handler.GetUserById(w, r, id)
	})(w, r)
}

func (a *ErrorAdapter) CreateUser(w http.ResponseWriter, r *http.Request) {
	server.WithErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return a.Handler.CreateUser(w, r)
	})(w, r)
}

func (a *ErrorAdapter) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	server.WithErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return a.Handler.UpdateUser(w, r, id)
	})(w, r)
}

func (a *ErrorAdapter) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	server.WithErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return a.Handler.DeleteUser(w, r, id)
	})(w, r)
}
