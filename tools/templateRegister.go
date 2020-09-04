package tools

import (
	"errors"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"sync"
)
var onceDo sync.Once
var templates map[string]*template.Template
func GetTemplates(e *echo.Echo) map[string]*template.Template {
	onceDo.Do(func() {

		templates = make(map[string]*template.Template)
		//templates["home.html"] = template.Must(template.ParseFiles("vp/view/home.html"))
		//templates["about.html"] = template.Must(template.ParseFiles("vp/view/about2.html")

		e.Renderer = &TemplateRegistry{
			templates: templates,
		}
	})
	return templates
}
// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, name, data)
}


