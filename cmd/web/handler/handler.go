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
func Guestbook(c echo.Context) error {
	time := time.Now()
	println("request received")
	return c.Render(http.StatusOK, "guestbook", time)
}
func Bex(c echo.Context) error {
	args := make([]int, 1)
	return c.Render(http.StatusOK, "dog", args)
}
