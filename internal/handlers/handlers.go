package handlers

import (
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/app/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}