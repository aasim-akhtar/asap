package views

import (
	"html/template"
	"path/filepath"
)

var (
	// @TODO use filepath.Join("views","layouts") with a trailing slash
	LayoutDir string = "views/layouts/"
	TemplateExt string = "*.html"
)

func NewView (layout string, files ...string) *View {

	files = append(files, layoutFiles()...)

	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: template,
		Layout : layout,
	}
}

// View struct contains pointer to a Template that produces a safe HTML document fragment.
type View struct {
	Template *template.Template
	Layout string
}

// layoutFiles returns a slice of strings representing
// the layout files used in our application.
func layoutFiles() []string {
	matchingFiles, err := filepath.Glob(LayoutDir + TemplateExt)
	if err != nil {
		panic(err)
	}
	return matchingFiles
}