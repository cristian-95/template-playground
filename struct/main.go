package main

import (
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Nome      string
	Sobrenome string
	Email     string
	Idade     int
}

func main() {
	http.HandleFunc("/", userHandler)
	log.Print("serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// usuario teste:
	User1 := User{
		Nome:      "João",
		Sobrenome: "Silva",
		Email:     "joao.silva@domain.com",
		Idade:     29,
	}
	tmpl, _ := template.ParseFiles("./public/user.html")
	tmpl.Execute(w, User1)
}
