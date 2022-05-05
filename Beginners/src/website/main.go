package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var funcMap = template.FuncMap{
	"Upper":   strings.ToUpper,
	"Multipl": Multipl,
}

var tpl *template.Template

type Page struct {
	Title string
	P     float32
}

func Multipl(p float32) (str string) {
	return fmt.Sprintf("%.2f", p*5)
}

var p = Page{
	Title: "Lothar",
	P:     23.6,
}
var err template.Error

func main() {
	tpl = template.New("index.html")
	tpl = tpl.Funcs(funcMap)
	tpl = template.Must(tpl.ParseFiles("templates/index.html"))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Server läuft Port 8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, p)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl2 := template.Must(template.ParseFiles("templates/login.templ.html"))
	tpl2.Execute(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tpl2 := template.Must(template.ParseFiles("templates/contact.templ.html"))
	tpl2.Execute(w, nil)
}
