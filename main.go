package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", nextWeek)

	log.Print("serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func nextWeek(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")

}
