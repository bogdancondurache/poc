package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"runtime"
)

//go:embed dist
var content embed.FS

func main() {
	http.HandleFunc("/index.html", httpHandler)
	http.Handle("/", fileHandler())
	http.ListenAndServe(":8080", nil)
}

func fileHandler() http.Handler {
	fsys := fs.FS(content)
	html, _ := fs.Sub(fsys, "dist")

	return http.FileServer(http.FS(html))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	fsys := fs.FS(content)
	htmlFs, err := fs.Sub(fsys, "dist")
	if err != nil {
		panic(err)
	}
	templ, err := template.ParseFS(htmlFs, "index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = templ.Execute(w, os)
	if err != nil {
		fmt.Println(err)
	}
}
