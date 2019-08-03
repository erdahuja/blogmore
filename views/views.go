package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var layoutDir = "layouts/"
var templateExt = ".gohtml"
var templateDir = "views/"

// View type containing template and error
type View struct {
	Err      error
	Template *template.Template
}

// New parses a template for html
func New(files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	tpl, err := template.ParseFiles(files...)
	return &View{
		Template: tpl,
		Err:      err,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(templateDir + layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// Render the view
func (v *View) Render(w http.ResponseWriter, lyt string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	if err := v.Template.ExecuteTemplate(w, lyt, data); err != nil {
		return (err)
	}
	return nil
}

// addTemplatePath prepends template dir directory to each file path
// eg: {home} would be updated to "views/home"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = templateDir + f
	}
}

// addTemplateExt prepends template dir directory to each file path
// eg: {home} would be updated to "home.gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + templateExt
	}
}
