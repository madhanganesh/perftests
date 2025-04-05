package utils

import (
	"errors"
	"net/http"
	"path/filepath"
	"text/template"
)

type TemplCache struct {
	cache map[string]*template.Template
}

func LoadTemplates() (*TemplCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		if name == "base" {
			continue
		}

		template, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		template, err = template.ParseFiles("./templates/base.html")
		if err != nil {
			return nil, err
		}

		template, err = template.ParseGlob("templates/partials/*.html")
		if err != nil {
			return nil, err
		}

		cache[name] = template
	}

	return &TemplCache{cache: cache}, nil
}

func (t *TemplCache) Render(w http.ResponseWriter, tmplName string, data any) error {
	template, ok := t.cache[tmplName]
	if !ok {
		return errors.New("template not found")
	}

	return template.Execute(w, data)
}
