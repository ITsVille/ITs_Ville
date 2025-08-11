package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Staattiset tiedostot (CSS, JS, kuvat)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Reitti pääsivulle
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		
		data := map[string]string{
			"Name": "Ville Nurmenniemi",
			"Role": "Software Developer & AI Enthusiast",
			"Description": "Hi, Welcome to my portfolio! Plaese say hello to my chat bot! :)",
		}
		
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}