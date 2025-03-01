package store_test

import (
	"github.com/stretchr/testify/assert"
	"goshort/internal/store"
	"testing"
)

func TestInMemoryURLStore_Set(t *testing.T) {
	t.Run("Returns true when the key doesn't exist", func(t *testing.T) {
		s := store.NewURLStore()
		hash, url := "hash", "url"

		result := s.Add(hash, url)

		assert.True(t, result)
	})

	t.Run("Returns false when the key already exists", func(t *testing.T) {
		s := store.NewURLStore()
		hash, url := "hash", "url"

		result := s.Add(hash, url)
		assert.True(t, result)

		result = s.Add(hash, url)
		assert.False(t, result)
	})
}

func TestInMemoryURLStore_Get(t *testing.T) {
	var cases = []struct {
		name, expectedURL string
		expectedResult    bool
	}{
		{
			"Returns an empty string when the URL doesn't exist in the store",
			"",
			false,
		},
		{
			"Gets the URL when it exists",
			"url",
			true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			s := store.NewURLStore()
			hash, expectedURL := "hash", test.expectedURL

			if expectedURL != "" {
				isSet := s.Add(hash, expectedURL)
				assert.True(t, isSet, "couldn't add the URL to the store")
			}

			url, result := s.Get(hash)
			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, expectedURL, url)
		})
	}
}
