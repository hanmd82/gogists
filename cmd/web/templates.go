package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/hanmd82/gogists/pkg/forms"
	"github.com/hanmd82/gogists/pkg/models"
)

type templateData struct {
	CurrentYear int
	Form        *forms.Form
	Gist        *models.Gist
	Gists       []*models.Gist
}

func formatDateTime(t time.Time) string {
	return t.Format(time.RFC822)
}

var functions = template.FuncMap{
	"formatDateTime": formatDateTime,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Get a slice of all the 'page' templates.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Create an empty template set, register functions in template.FuncMap, then parse the file
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add any 'layout' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add any 'partial' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
