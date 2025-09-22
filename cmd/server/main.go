package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/handlers"
	"url-shortener/internal/storage"
)

func main() {
	fmt.Println("Starting URL Shortener...")

	store := storage.NewMemorySrorage()

	handlers := handlers.NewHandlers(store)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/shorten", handlers.Shorten)
    http.HandleFunc("/r/", handlers.Redirect)

	log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}