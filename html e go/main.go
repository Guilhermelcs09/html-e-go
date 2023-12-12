package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	nome  string
	email string
}

var templates template.Template

func main() {

	u := usuario{
		nome:  "Joao",
		email: "joaozin@gmail.com",
	}

	templates = *template.Must(template.ParseGlob("*html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
