package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var (
	t map[string]*template.Template
	b bytes.Buffer
)

func init() {
	t = make(map[string]*template.Template)
	t["user.html"] = template.Must(template.ParseFiles("base.html", "user.html"))
	t["page.html"] = template.Must(template.ParseFiles("base.html", "page.html"))
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle",
	}
	err := t["page.html"].ExecuteTemplate(&b, "base", p)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		return
	}
	b.WriteTo(w)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := User{
		Username: "swordsmith",
		Name:     "Inigo Montoya",
	}
	err := t["user.html"].ExecuteTemplate(&b, "base", u)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		return
	}
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
