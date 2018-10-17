package main

import (
	"fmt"
	"net/http"
	"text/template"
	"github.com/GeertJohan/go.rice"
)

var (
	templateFilesBox *rice.Box
	libFilesBox      *rice.Box
)

func main() {
	templateFilesBox = rice.MustFindBox("dist")
	libFilesBox = rice.MustFindBox("dist/static")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index.html", indexPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(libFilesBox.HTTPBox())))
	fmt.Println("listen in 127.0.0.1:8888")
	http.ListenAndServe(":8888", nil)
}

// rootHandler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// 跳转到 /index.html
		w.Header().Set("Location", "/index.html")
		http.Error(w, "", 301)
	} else {
		http.Error(w, "404 Not Found", 404)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	contents, err := templateFilesBox.String("index.html")
	if err != nil {
		return
	}
	tmpl, err := template.New("index.html").Parse(contents)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
