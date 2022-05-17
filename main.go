package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", myTemplateHandler)
	fmt.Println(nextWeek().String())
	log.Print("serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myTemplateHandler(w http.ResponseWriter, r *http.Request) {
	releaseDate := nextWeek()
	Date := releaseDate.String()

	tmpl, _ := template.ParseFiles("./public/index.html")
	tmpl.Execute(w, Date)
}

func nextWeek() time.Time {
	rd := time.Now()
	return rd.Add(7 * 24 * time.Hour)

}
