package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.html", "head.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle",
	}

	var b bytes.Buffer

	err := t.ExecuteTemplate(&b, "index.html", p)

	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		return
	}

	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
