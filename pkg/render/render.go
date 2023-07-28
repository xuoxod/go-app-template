package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/xuoxod/go-app-template/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Printf("\n\tCurrent page is: %v\n", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		// Find base layout and parse
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			fmt.Println("Layout match? ", matches)
			return myCache, err
		}

		if len(matches) > 0 {
			fmt.Println("Matched Layout: ", matches)
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}
