package views
import (
"html/template"
"path/filepath"
	"net/http"
)
var LayoutDir string = "views/layouts"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout: layout,
	}
}
type View struct {
	Template *template.Template
	Layout string

}
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}