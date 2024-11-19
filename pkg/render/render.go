package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/FloMatt/SimpleGoWebsite/pkg/config"
	"github.com/FloMatt/SimpleGoWebsite/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// CreateTemplateCache reads all the templates from the templates directory and caches them
// RenderTemplate writes the rendered template to the http.ResponseWriter.
// It retrieves the template from the app's TemplateCache based on the provided template name (tmpl).
// If the template is not found in the cache, it logs a fatal error.
// The template is then executed into a bytes.Buffer, and the buffer's content is written to the http.ResponseWriter.
// If an error occurs during writing or executing the template, it logs the error and returns.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("couldn't find template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render the template

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing to response: ", err)

	}

	/*parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	  err := parsedTemplate.Execute(w, nil)
	  if err != nil {
	      fmt.Println("Error parsing template: ", err)
	      return
	  }*/
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

/*var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if already we have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		//need to create a new template
		log.Println("creating new tempate")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("Error creating template cache: ", err)
			return
		}
		//use the newly created template
		log.Println("created new tempate")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("Error creating template cache: ", err)
			return
		}
		//use the newly created template
		log.Println("created new tempate")
	} else {
		//use the template from our cache
		log.Println("using cached tempate")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	//add the template to our cache
	tc[t] = tmpl
	return nil
}*/
