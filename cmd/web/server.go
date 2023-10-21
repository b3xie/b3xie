package main

import (
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func Hello(c echo.Context) error {
	time := time.Now()
	return c.Render(http.StatusOK, "hello", time)
}
func Dog(c echo.Context) error {
	time := time.Now()
	println("request received")
	return c.Render(http.StatusOK, "dog", time)
}
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob(
			"templates/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/hello", Hello)
	e.Static("/dist", "dist")
	e.GET("/dog", Dog)
	e.File("/borzoi", "static/borzoi.jpeg")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1328"))
}
