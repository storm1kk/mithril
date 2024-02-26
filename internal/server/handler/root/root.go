package root

import (
	"html/template"
	"net/http"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/app/index.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			return
		}
	})
}
