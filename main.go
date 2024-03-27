package main

import (
	"net/http"

	"scroll2top.com/golangwebappbnb/config"
	"scroll2top.com/golangwebappbnb/handlers"
)

func main() {
	var app config.AppConfig

	app.SomeGlobalValue = "HelloGlobalValue"

	handlers.SetConfig(&app)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	http.ListenAndServe(":8080", nil)
}
