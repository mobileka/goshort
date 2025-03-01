package ui_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"goshort/internal/ui"
)

const (
	templatesPath = "testdata/templates"
)

func TestLoadTemplates(t *testing.T) {
	t.Run("when templatesPath is correct", func(t *testing.T) {
		expectedKeys := []string{"index.html", "error.html", "result.html"}

		result, err := ui.LoadTemplates(templatesPath)
		assert.NoError(t, err)

		for _, expectedKey := range expectedKeys {
			assert.Contains(t, result, expectedKey)

			// see testdata/templates/#{expectedKey}
			expectedValue := "<title>" + expectedKey + "</title>"

			// Render the template to a string to compare to the expected result
			var tplBuf bytes.Buffer
			err := result[expectedKey].Execute(&tplBuf, nil)
			assert.Nil(t, err)

			renderResult := tplBuf.String()
			assert.Equal(t, expectedValue, renderResult)
		}
	})

	t.Run("when templatesPath doesn't exist", func(t *testing.T) {
		result, err := ui.LoadTemplates("some_weird_path")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestMustLoadTemplates(t *testing.T) {
	t.Run("when templatesPath is correct", func(t *testing.T) {
		notPanicFunc := func() { ui.MustLoadTemplates(templatesPath) }
		assert.NotPanics(t, notPanicFunc)
	})

	t.Run("when templatesPath doesn't exist", func(t *testing.T) {
		panicFunc := func() { ui.MustLoadTemplates("some_weird_path") }
		assert.Panics(t, panicFunc)
	})
}
