package utils

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

var dirs = []string{
	"views/*.html",
	"views/layout/*.html",
	"views/pages/*.html",
}

type Renderer struct {
	template *template.Template
}

func NewRenderer() *Renderer {
	tpl := new(Renderer)
	tpl.ReloadTemplates()
	return tpl
}

func (t *Renderer) ReloadTemplates() {
	files := []string{}
	for _, dir := range dirs {
		ff, err := filepath.Glob(dir)
		if err != nil {
			panic(err)
		}
		files = append(files, ff...)
	}

	t.template = template.Must(template.ParseFiles(files...))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.template.ExecuteTemplate(w, name, data)
}
