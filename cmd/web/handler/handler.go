package handler

import (
	"encoding/csv"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type Post struct {
	Name    string
	Message string
}

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
	return c.Render(http.StatusOK, "bex", args)
}
func AddGuestbookEntry(c echo.Context) error {
	name := c.FormValue("guestName")
	text := c.FormValue("guestText")
	post := [][]string{
		{name, text},
	}
	file, err := os.OpenFile("internal/posts/posts.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = csv.NewWriter(file).WriteAll(post)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "guestbookpost", name)
}
