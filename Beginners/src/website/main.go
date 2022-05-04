package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/hello", helloHandleFunc)
	fmt.Println("Server läuft Port 8080")
	http.ListenAndServe(":8080", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.templ.html")
	if err != nil {
		fmt.Fprint(w, err.Error(), "helloHandleFunc")
	}
	tmpl.Execute(w, Page{Title: "Lothar"})
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.templ.html")
	if err != nil {
		fmt.Fprint(w, err.Error(), "helloHandleFunc")
	}
	tmpl.Execute(w, Page{Title: "Lothar"})
}