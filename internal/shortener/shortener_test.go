package shortener_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goshort/internal/storetest"

	"goshort/internal/shortener"
)

func newShortener(url string, result bool) *shortener.Shortener {
	store := storetest.NewStoreMock(url, result)
	return shortener.NewShortener(store)
}

func TestShortener_Shorten(t *testing.T) {
	t.Run("Shortens the URL", func(t *testing.T) {
		url := "url"
		success := true
		s := newShortener(url, success)

		result, err := s.Shorten(url)

		assert.Nil(t, err)
		assert.IsType(t, "", result)
	})

	t.Run("Returns an error when tries too many times and fails", func(t *testing.T) {
		url := "url"
		success := false
		s := newShortener(url, success)

		result, err := s.Shorten(url)

		assert.Equal(t, "", result)
		assert.Error(t, err, result)
	})
}

func TestShortener_Expand(t *testing.T) {
	t.Run("Expands when the URL exists", func(t *testing.T) {
		expectedURL := "url"
		s := newShortener(expectedURL, true)

		url, ok := s.Expand("hash")

		assert.True(t, ok)
		assert.Equal(t, expectedURL, url)
	})

	t.Run("Returns an empty result if the URL doesn't exist", func(t *testing.T) {
		expectedURL := ""
		s := newShortener(expectedURL, false)

		url, ok := s.Expand("hash")

		assert.False(t, ok)
		assert.Equal(t, expectedURL, url)
	})
}
