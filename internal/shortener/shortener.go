package shortener

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"urlshortener/internal/store"
)

const (
	hashLength         = 6
	maxShortenAttempts = 100
)

// Shortener provides URL shortening functionality
type Shortener struct {
	store store.URLStore
}

// NewShortener creates a new URL shortener service
func NewShortener(store store.URLStore) *Shortener {
	return &Shortener{
		store: store,
	}
}

// Shorten creates a hash for the given URL
func (s *Shortener) Shorten(url string) (string, error) {
	// Generate a hash and ensure it's unique
	attempts := 0
	for {
		attempts += 1
		if attempts > maxShortenAttempts {
			return "", fmt.Errorf("failed to shorten after %d attempts", attempts)
		}

		hash := generateHash(hashLength)
		if s.store.Set(hash, url) {
			return hash, nil
		}
	}
}

// Expand retrieves the original URL for a hash
func (s *Shortener) Expand(hash string) (string, bool) {
	return s.store.Get(hash)
}

// generateHash creates a random alphanumeric hash of the given length
func generateHash(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hash := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		hash[i] = charset[randomIndex.Int64()]
	}

	return string(hash)
}
