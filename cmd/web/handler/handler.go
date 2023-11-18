package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type Post struct {
	Name    string `json:"Name"`
	Message string `json:"Message"`
}

func ErrorHandler(err error, c echo.Context) {
	a, _ := err.(*echo.HTTPError)
	c.Render(http.StatusNotFound, "error", a)
}
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
	fmt.Println("post")
	if c.FormValue("guestName") == "" || c.FormValue("guestText") == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	var received = []Post{
		{
			Name:    c.FormValue("guestName"),
			Message: c.FormValue("guestText"),
		},
	}
	file, err := os.OpenFile("internal/posts/posts.json", os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var posts []Post
	json.Unmarshal(data, &posts)
	result := append(received, posts...)
	write, _ := json.MarshalIndent(result, "", "	")
	os.WriteFile("internal/posts/posts.json", write, 0664)
	if err != nil {
		panic(err)
	}

	return c.Render(http.StatusOK, "guestbookpost", result)
}
func GetGuestbookentries(c echo.Context) error {
	file, err := os.OpenFile("internal/posts/posts.json", os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var posts []Post
	json.Unmarshal(data, &posts)
	return c.Render(http.StatusOK, "guestbookpost", posts)
}
