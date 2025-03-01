package ui

import (
	"html/template"
	"os"
	"path/filepath"
)

// LoadTemplates loads HTML templates from the specified path.
// Returns a map of template names to their corresponding *template.Template objects.
func LoadTemplates(templatesPath string) (map[string]*template.Template, error) {
	files, err := os.ReadDir(templatesPath)
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(templatesPath, file.Name())
		templates[file.Name()] = template.Must(template.ParseFiles(path))
	}

	return templates, nil
}

// MustLoadTemplates loads HTML templates from the specified path.
// Panics if an error occurs during the loading process.
// Returns a map of template names to their corresponding *template.Template objects.
func MustLoadTemplates(templatesPath string) map[string]*template.Template {
	templates, err := LoadTemplates(templatesPath)

	if err != nil {
		panic(err)
	}

	return templates
}
