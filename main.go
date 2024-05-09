package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	films := []Film{
		{Title: "Casablanca", Director: "Michael Curtiz"},
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "Raging Bull", Director: "Martin Scorsese"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, map[string][]Film{"Films": films})
	})

	http.HandleFunc("/form.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		films = append(films, Film{Title: title, Director: director})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
