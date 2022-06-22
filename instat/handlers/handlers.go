package handlers

import (
	"html/template"
	"net/http"
)

var temp *template.Template

func init() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	temp = template.Must(template.ParseGlob("template/*.html"))
}

func RunIndex(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "contact.html", nil)
}
