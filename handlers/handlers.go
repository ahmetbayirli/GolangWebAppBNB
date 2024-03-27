package handlers

import (
	"net/http"

	"scroll2top.com/golangwebappbnb/config"
	"scroll2top.com/golangwebappbnb/models"
)

var appConfig config.AppConfig

func SetConfig(globalConf *config.AppConfig){
	appConfig = *globalConf
}

func HomeHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) {
	RenderTemplate(responseWriter, "home.page.tmpl", &models.PageData{})
}

func AboutHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = appConfig.SomeGlobalValue
	RenderTemplate(responseWriter, "about.page.tmpl", &models.PageData{StrMap: strMap})
}