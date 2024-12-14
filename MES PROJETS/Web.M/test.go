package main

import (
	"html/template"
	"net/http"
)

const port = ":2003"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/tchewa", About)
	http.ListenAndServe(port, nil)
}
