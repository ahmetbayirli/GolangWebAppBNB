package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",  func(responseWriter http.ResponseWriter, requestPointer *http.Request){
		fmt.Fprintf(responseWriter , "Hello World")
	})

	http.ListenAndServe(":8080", nil)
}