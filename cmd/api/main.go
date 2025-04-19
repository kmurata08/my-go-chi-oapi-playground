package main

import (
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/server"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/user"

	"log"
	"net/http"
)

func main() {
	r := server.NewRouter()

	userService := user.NewService()
	userHandler := user.NewHandler(userService)
	userHandler.Register(r)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
