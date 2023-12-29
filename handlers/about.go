package handlers

import (
	"html/template"
	"net/http"
)

func (p *Product) GetAbout(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(Base, About, Footer))
	tmpl.ExecuteTemplate(rw, "about.html", nil)

}
