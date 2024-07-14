package sava

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("templates/template.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// define the data to return
	t.ExecuteTemplate(w, "template.html", "nil")
}
