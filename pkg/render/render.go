package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ikapfw/bnb-booking-web-app/pkg/config"
	"github.com/ikapfw/bnb-booking-web-app/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

// sets config for template package
func NewTemplate(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

// render template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
	// create template cache
	// tc, err := CreateTemplateCache()

	// if err != nil{
	// 	log.Fatal(err)
	// }
	var tc map[string]*template.Template
	
	// get the template cache from app config
	if app.UseCache{
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	// err = t.Execute(buf, nil)
	_ = t.Execute(buf, td)

	// if err != nil {
	// 	log.Println()
	// }

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		// log.Println(err)
		log.Println("Error writing template to browser:", err)
	}
	
	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	// err := parsedTemplate.Execute(w, nil)

	// if err != nil{
	// 	fmt.Println("error parsing template:", err)
	// 	return
	// }
}

// create a complex template cache
func CreateTemplateCache() (map[string]*template.Template, error){
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil{
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil{
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil{
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// simple cache render template
// func RenderTemplate(w http.ResponseWriter, t string){
// 	var tmpl *template.Template
// 	var err error

	// check to see if the template is already in the cache
	// _, inMap := tc[t]

	// if !inMap{
		// need to create template
	// 	log.Println("creating template and add to cache")

	// 	err = createTemplateCache(t)

	// 	if err != nil{
	// 		log.Println(err)
	// 	}
	// } else {
		// template already exist in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// }

// create a simple template cache
// func createTemplateCache(t string) error{
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

	// parse the template
	// tmpl, err := template.ParseFiles(templates...)

	// if err != nil{
	// 	return err	
	// }

	// add template to cache (map)
// 	tc[t] = tmpl

// 	return nil
// }