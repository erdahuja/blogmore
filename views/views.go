package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var layoutDir = "views/layouts/"
var templateExt = ".gohtml"

// View type containing template and error
type View struct {
	Err      error
	Template *template.Template
}

// New parses a template for html
func New(files ...string) *View {
	files = append(files, layoutFiles()...)
	tpl, err := template.ParseFiles(files...)
	return &View{
		Template: tpl,
		Err:      err,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// Render the view
func (v *View) Render(w http.ResponseWriter, lyt string, data interface{}) error {
	if err := v.Template.ExecuteTemplate(w, lyt, nil); err != nil {
		return (err)
	}
	return nil
}
