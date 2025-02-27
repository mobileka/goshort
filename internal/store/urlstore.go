package store

import (
	"sync"
)

// URLStore defines the interface for URL storage
type URLStore interface {
	Get(hash string) (string, bool)
	Set(hash, url string) bool
}

// URLMap implements the URLStore interface with an in-memory map
type URLMap struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewURLMap creates a new URL mapping store
func NewURLMap() *URLMap {
	return &URLMap{
		store: make(map[string]string),
	}
}

// Get retrieves a URL by its hash
func (u *URLMap) Get(hash string) (string, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	url, exists := u.store[hash]
	return url, exists
}

// Set stores a URL with the given hash
// Returns false if the hash already exists
func (u *URLMap) Set(hash, url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exists := u.store[hash]; exists {
		return false
	}

	u.store[hash] = url
	return true
}
