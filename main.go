package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", myTemplateHandler)
	log.Print("serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myTemplateHandler(w http.ResponseWriter, r *http.Request) {
	date := releaseDate()
	tmpl, _ := template.ParseFiles("./public/index.html")
	tmpl.Execute(w, date)
}

func releaseDate() string {
	rd := time.Now()
	rd = rd.AddDate(0, 0, 7)
	return rd.Format("02/01/2006")
}
