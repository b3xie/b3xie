package main

import (
	"b3xie/cmd/web/handler"
	"b3xie/cmd/web/mdparser"
	"html/template"
	"io"
	"time"

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
	e.Static("/dist", "dist")
	e.GET("/guestbook", handler.Guestbook)
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 2, Burst: 1, ExpiresIn: 1 * time.Hour},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
	}
	e.GET("/", handler.Index)
	e.GET("/redirect", handler.HtmxRefresh)
	e.GET("/bex", handler.Bex)
	e.POST("/guestbook/add", handler.AddGuestbookEntry, middleware.RateLimiterWithConfig(config))
	e.GET("guestbook/get", handler.GetGuestbookentries)
	e.DELETE("guestbook/delete", handler.GuestbookDelete)
	e.HTTPErrorHandler = handler.ErrorHandler
	mdparser.ParseNewFiles()

	e.Logger.Fatal(e.Start(":8080"))

}
