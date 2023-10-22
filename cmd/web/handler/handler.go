package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NotFound() {}
func Hello(c echo.Context) error {
	time := time.Now()
	return c.Render(http.StatusOK, "index", time)
}
func Dog(c echo.Context) error {
	time := time.Now()
	println("request received")
	return c.Render(http.StatusOK, "dog", time)
}
