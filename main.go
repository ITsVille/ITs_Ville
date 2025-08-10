package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Staattiset tiedostot (CSS, JS, kuvat)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Reitti pääsivulle
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := map[string]string{
			"Name": "Ville Nurmenniemi",
			"Role": "Software Developer & AI Enthusiast",
		}
		tmpl.Execute(w, data)
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
