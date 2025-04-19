package server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/errors"
	"net/http"
)

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func WithErrorHandler(h ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	var status int
	var response any

	switch e := err.(type) {
	case *errors.APIError:
		status = e.StatusCode
		response = e
	default:
		status = http.StatusInternalServerError
		response = &errors.APIError{
			StatusCode: status,
			Code:       "internal_server_error",
			Message:    "An unexpected error occurred",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	// 共通ミドルウェアの設定があればここに書く
	// r.Use(middleware.Logger)
	return r
}
