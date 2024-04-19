package handlers

import (
	"log"
	"net/http"

	"scroll2top.com/golangwebappbnb/config"
	"scroll2top.com/golangwebappbnb/models"
)

var appConfig config.AppConfig

func SetConfig(globalConf *config.AppConfig) {
	appConfig = *globalConf
}

func HomeHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) {
	RenderTemplate(responseWriter, "home.page.tmpl", &models.PageData{})
}

func LoginHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) {
	err := requestPointer.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	strMap := make(map[string]string)
	strMap["user_name"] = requestPointer.Form.Get("user_name")
	strMap["password"] = requestPointer.Form.Get("password")

	RenderTemplate(responseWriter, "home.page.tmpl", &models.PageData{StrMap: strMap})
}

func AboutHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = appConfig.SomeGlobalValue
	RenderTemplate(responseWriter, "about.page.tmpl", &models.PageData{StrMap: strMap})
}
