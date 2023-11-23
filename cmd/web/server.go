package main

import (
	"b3xie/cmd/web/handler"
	"b3xie/cmd/web/mdparser"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3)))
	e.Static("/dist", "dist")
	e.GET("/guestbook", handler.Guestbook)
	e.GET("/", handler.Hello)
	e.GET("/bex", handler.Bex)
	e.POST("/guestbook/add", handler.AddGuestbookEntry)
	e.GET("guestbook/get", handler.GetGuestbookentries)
	e.HTTPErrorHandler = handler.ErrorHandler
	mdparser.ParseNewFiles()
	
	e.Logger.Fatal(e.Start(":1329"))

}
