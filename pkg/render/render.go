package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//get requested template from cache
	t, ok 

	//render the template

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all the files in named .page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through all the files ending in *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
	}
	matches, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return myCache, err
	}
	if len(matches) > 0 {
		{
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache{name} = ts
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
