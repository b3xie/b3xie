package main

import (
	"b3xie/cmd/web/handler"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob(
			"templates/*.html")),
	}
	e.Renderer = t
	e.Static("/dist", "dist")
	e.GET("/dog", handler.Dog)
	e.GET("/", handler.Hello)
	e.Logger.Fatal(e.Start(":1329"))
}
