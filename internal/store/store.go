package store

import (
	"sync"
)

// URLStore defines the interface for URL storage
type URLStore interface {
	Get(hash string) (string, bool)
	Add(hash, url string) bool
}

// InMemoryURLStore implements the URLStore interface with an in-memory sync.Map
type InMemoryURLStore struct {
	store sync.Map
}

// NewURLStore creates a new URLStore instance
func NewURLStore() URLStore {
	return &InMemoryURLStore{}
}

// Get retrieves a URL by its hash
func (s *InMemoryURLStore) Get(hash string) (string, bool) {
	url, exists := s.store.Load(hash)

	if !exists {
		return "", false
	}

	return url.(string), exists
}

// Set stores a URL with the given hash
// Returns false if the hash already exists
func (s *InMemoryURLStore) Add(hash, url string) bool {
	_, exists := s.store.LoadOrStore(hash, url)

	return !exists
}
