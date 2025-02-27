package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"urlshortener/internal/handler"
	"urlshortener/internal/shortener"
	"urlshortener/internal/store"
)

const (
	baseUrl = "http://localhost:8080/"
)

func main() {
	// Initialize the URL store
	urlStore := store.NewURLMap()

	// Initialize the shortener service
	shortenerService := shortener.NewShortener(urlStore)

	// Path to templates
	templatesPath := filepath.Join("ui", "templates")

	// Initialize the HTTP handler
	h := handler.NewHandler(shortenerService, baseUrl, templatesPath)

	registerRoutes(h)

	// Start the server
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerRoutes(h *handler.Handler) {
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/shorten", h.ShortenURL)
}
