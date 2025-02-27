package handler

import (
	"goshort/internal/shortener"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

// Handler manages HTTP request handling
type Handler struct {
	shortener     *shortener.Shortener
	baseURL       string
	templates     map[string]*template.Template
	templatesPath string
}

// TemplateData holds data for HTML templates
type TemplateData struct {
	OriginalURL  string
	ShortURL     string
	ErrorMessage string
}

// NewHandler creates a new HTTP handler
func NewHandler(shortener *shortener.Shortener, baseURL string, templatesPath string) *Handler {
	// Create a new handler
	h := &Handler{
		shortener:     shortener,
		baseURL:       baseURL,
		templates:     make(map[string]*template.Template),
		templatesPath: templatesPath,
	}

	// Load templates
	h.loadTemplates()

	return h
}

// loadTemplates loads all HTML templates
func (h *Handler) loadTemplates() {
	// Define template names
	templates := []string{"index.html", "result.html", "error.html"}

	// Load each template
	for _, tmpl := range templates {
		path := filepath.Join(h.templatesPath, tmpl)
		h.templates[tmpl] = template.Must(template.ParseFiles(path))
	}
}

// renderTemplate renders a template with the given data
func (h *Handler) renderTemplate(w http.ResponseWriter, tmpl string, data TemplateData) {
	// Get template from the cache
	t, exists := h.templates[tmpl]
	if !exists {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Home renders the homepage
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// Only handle the root path
	if r.URL.Path != "/" {
		h.handleRedirect(w, r)
		return
	}

	// Render the index template
	h.renderTemplate(w, "index.html", TemplateData{})
}

// ShortenURL handles form submissions to create short URLs
func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse the form
	err := r.ParseForm()
	if err != nil {
		h.renderTemplate(w, "error.html", TemplateData{
			ErrorMessage: "Invalid form submission",
		})
		return
	}

	// Get the URL from the form
	url := r.FormValue("url")

	// Validate the URL
	if url == "" {
		h.renderTemplate(w, "error.html", TemplateData{
			ErrorMessage: "URL is required",
		})
		return
	}

	// Add https:// prefix if missing
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	// Shorten the URL
	hash, err := h.shortener.Shorten(url)
	if err != nil {
		// Render the result template
		h.renderTemplate(w, "error.html", TemplateData{
			ErrorMessage: "Cannot shorten the URL: too many collisions",
		})
		return
	}

	shortURL := h.baseURL + hash

	// Render the result template
	h.renderTemplate(w, "result.html", TemplateData{
		OriginalURL: url,
		ShortURL:    shortURL,
	})
}

// handleRedirect handles redirection of short URLs to their original destinations
func (h *Handler) handleRedirect(w http.ResponseWriter, r *http.Request) {
	// Skip the leading slash to get the hash
	hash := r.URL.Path[1:]

	// Look up the original URL
	originalURL, exists := h.shortener.Expand(hash)
	if !exists {
		h.renderTemplate(w, "error.html", TemplateData{
			ErrorMessage: "Short URL not found",
		})
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
