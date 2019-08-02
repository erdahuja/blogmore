package views

import "text/template"

type View struct {
	Err      error
	Template *template.Template
}

// New parses a template for html
func New(files ...string) *View {
	files = append(files, "views/layouts/header.gohtml", "views/layouts/footer.gohtml")
	tpl, err := template.ParseFiles(files...)
	return &View{
		Template: tpl,
		Err:      err,
	}
}
