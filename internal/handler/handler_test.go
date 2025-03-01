// handler_test/handler_test.go
package handler_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"goshort/internal/shortener"
	"goshort/internal/store"
	"goshort/internal/storetest"

	"goshort/internal/handler"
)

const (
	templatesPath = "testdata/templates"
	baseUrl       = "http://example.com"
)

func newHandler(store store.URLStore) *handler.Handler {
	templates := make(map[string]*template.Template)
	for _, tpl := range []string{"index.html", "error.html", "result.html"} {
		path := filepath.Join(templatesPath, tpl)
		templates[tpl] = template.Must(template.ParseFiles(path))
	}
	sh := shortener.NewShortener(store)
	return handler.NewHandler(sh, baseUrl, templates)
}

func TestHome(t *testing.T) {
	h := newHandler(storetest.NewSucceedingStoreMock(baseUrl))
	expectedBody := "<title>index.html</title>"

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	h.Home(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedBody, rr.Body.String())
}

func TestShortenURL(t *testing.T) {
	t.Run("when everything goes smoothly", func(t *testing.T) {
		h := newHandler(storetest.NewSucceedingStoreMock(baseUrl))
		expectedBody := "<title>result.html</title>"

		formData := "url=http://example.com"
		req, err := http.NewRequest("POST", "/shorten", strings.NewReader(formData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		h.ShortenURL(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedBody, rr.Body.String())
	})

	t.Run("when storage fails", func(t *testing.T) {
		h := newHandler(storetest.NewFailingStoreMock(baseUrl))
		expectedBody := "<title>error.html</title>"

		formData := "url=http://example.com"
		req, err := http.NewRequest("POST", "/shorten", strings.NewReader(formData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		h.ShortenURL(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedBody, rr.Body.String())
	})

	t.Run("when HTTP method is not POST", func(t *testing.T) {
		// TODO: implement me
		// http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	t.Run("when form data is invalid", func(t *testing.T) {
		// TODO: implement me
	})

	t.Run("when can't parse the form data", func(t *testing.T) {
		// TODO: implement me
	})
}

func TestHandleRedirect(t *testing.T) {
	expectedURL := "http://someurl.com"
	h := newHandler(storetest.NewSucceedingStoreMock(expectedURL))

	req, err := http.NewRequest("GET", "/abc123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	h.Home(rr, req)

	assert.Equal(t, http.StatusMovedPermanently, rr.Code)
	assert.Equal(t, expectedURL, rr.Header().Get("Location"))
}
