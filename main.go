package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/index.html", httpHandler)
	http.HandleFunc("/", fileHandler)
	http.ListenAndServe(":8080", nil)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := strings.TrimPrefix(r.URL.Path, "/")
	fileBytes, err := ioutil.ReadFile("./dist/" + fileName)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(fileBytes)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	templ, err := template.ParseFiles("./dist/index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = templ.Execute(w, os)
	if err != nil {
		fmt.Println(err)
	}
}
