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
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	//handler is a file server that reads from static dir. 
	http.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":8080", nil)
}
