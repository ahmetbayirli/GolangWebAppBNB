package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) {
	renderTemplate(responseWriter, "home.page.tmpl")
}

func renderTemplate(responseWriter http.ResponseWriter, templateName string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + templateName)

	handleError(err)

	err = parsedTemplate.Execute(responseWriter, nil) // second argument no Data

	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", nil)
}
