package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Context-Type", "text/html: charset=utf-8")
	err := t.htmlTmpl.Execute(w, data)

	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template failed %w", err)
	}
	return Template{htmlTmpl: tpl}, nil
}
