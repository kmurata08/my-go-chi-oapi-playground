package main

import (
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/server"

	"log"
	"net/http"
)

func main() {
	// ルーターの作成
	r := server.NewRouter()

	// サーバー起動
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
