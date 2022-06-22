package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	// http.Handle("/static/", http.StripPrefix("/static", fs))

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/blog", blogHandler)
	router.HandleFunc("/blogsingle", blogSingleHandler)

	http.ListenAndServe(":8080", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "blog.html", nil)
}

func blogSingleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "blog-single.html", nil)
}
