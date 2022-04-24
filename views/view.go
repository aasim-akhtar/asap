package views

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	LayoutDir   string = filepath.Join("views", "layouts") + string(os.PathSeparator)
	TemplateExt string = "*.html"
)

func NewView(layout string, files ...string) *View {

	files = append(files, layoutFiles()...)

	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: template,
		Layout:   layout,
	}
}

// View struct contains pointer to a Template that produces a safe HTML document fragment.
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
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
