package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"scroll2top.com/golangwebappbnb/models"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, pd *models.PageData) {
	var tmpl *template.Template
	var err error

	_, inMap := tmplCache[t]

	if !inMap {
		err = makeTemplateCache(t)
		HandleError(err)
	} else {
		fmt.Println("Template in cache !! delete me on prod")
	}

	tmpl = tmplCache[t]
	err = tmpl.Execute(w, pd)
	HandleError(err)
}

func makeTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tmplCache[t] = tmpl
	return nil
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}


