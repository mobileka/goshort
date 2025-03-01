package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"goshort/internal/handler"
	"goshort/internal/shortener"
	"goshort/internal/store"
	"goshort/internal/ui"
)

const (
	port    = "8080"
	baseUrl = "http://localhost:" + port + "/"
)

func main() {
	// Initialize the URL store
	urlStore := store.NewURLStore()

	// Initialize the shortener service
	shortenerService := shortener.NewShortener(urlStore)

	// Path to templates
	templatesPath := filepath.Join("ui", "templates")
	templates := ui.MustLoadTemplates(templatesPath)

	// Initialize the HTTP handler
	h := handler.NewHandler(shortenerService, baseUrl, templates)

	registerRoutes(h)

	// Start the server
	fmt.Println("Server starting on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func registerRoutes(h *handler.Handler) {
	http.HandleFunc("/shorten", h.ShortenURL)
	http.HandleFunc("/", h.Home)
}
