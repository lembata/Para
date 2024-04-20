package api

import (
	"io"
	//"net/http"
	"html/template"
)


type Templates struct {
	templates *template.Template
}

// func (t *Templates) Render(w http.ResponseWriter, name string, data interface{}) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func (t *Templates) LoadTemplates() error {

	temp, err := template.ParseGlob("ui/views/*.html")

	if err != nil {
		return err
	}

	t.templates = template.Must(temp, err)

	return nil
}
