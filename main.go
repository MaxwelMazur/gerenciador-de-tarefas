package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	http.HandleFunc("/index", Index)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	println("É mais de oito mil!")
	http.ListenAndServe(":8001", nil)
}
