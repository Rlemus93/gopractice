package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Director: "Michael Curtiz"},
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Raging Bull", Director: "Martin Scorsese"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='py-4 text-xl font-semibold text-gray-800'>%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
