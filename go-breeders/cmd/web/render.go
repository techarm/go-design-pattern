package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, name string, td *templateData) {
	var tmpl *template.Template

	// if we are using the template cache, try to get the tempalte from our map, stored int the receiver
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[name]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(name)
		if err != nil {
			log.Println("error building template: ", err)
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		log.Println("building template from disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, name, td); err != nil {
		log.Println("error execution template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplateFromDisk(name string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", name),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[name] = tmpl
	return tmpl, nil
}
