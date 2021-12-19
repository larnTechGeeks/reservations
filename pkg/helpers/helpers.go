package helpers

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/larnTechGeeks/reservations/pkg/config"
	"github.com/larnTechGeeks/reservations/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func AddAppConfig(a *config.AppConfig) {
	app = a
}

func AddGlobalData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error
	if app.DEBUG {
		tc, err = BuildTC()
		if err != nil {
			log.Println(err)
		}
	} else {
		tc = app.TCache

		if err != nil {
			log.Println(err)
		}
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Println("Cannot find the given template")
	} else {
		// build a buffer
		buf := new(bytes.Buffer)
		// add global context data before execute
		td = AddGlobalData(td)
		t.Execute(buf, td)

		//write buffer to rw
		buf.WriteTo(rw)
	}
}

// BuildTC creates system wise template cache
func BuildTC() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	// capture all page templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// loop through the pages and parse their layouts
	for _, page := range pages {
		// get the page name -- used as key on TS
		name := filepath.Base(page)
		// build a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// check any layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// try and match the layout with the page
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts

	}

	return myCache, nil

}
