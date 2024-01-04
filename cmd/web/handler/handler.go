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
	Id      string
}

func HtmxRefresh(c echo.Context) error {
	c.Response().Header().Set("HX-Location", "/guestbook")
	return c.NoContent(http.StatusOK)
}
func ErrorHandler(err error, c echo.Context) {
	a, _ := err.(*echo.HTTPError)
	c.Render(http.StatusNotFound, "error", a)
}
func Index(c echo.Context) error {
	time := time.Now()
	return c.Render(http.StatusOK, "index", time)
}
func Guestbook(c echo.Context) error {
	return c.Render(http.StatusOK, "guestbook", "")
}
func Bex(c echo.Context) error {
	args := make([]int, 1)
	return c.Render(http.StatusOK, "bex", args)
}
func AddGuestbookEntry(c echo.Context) error {
	if c.FormValue("guestName") == "" || c.FormValue("guestText") == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	if len(c.FormValue("guestName")) > 50 || len(c.FormValue("guestText")) > 100 {
		return c.Render(http.StatusOK, "alert", "Too long! Your name should be <20 char and your message <100 char")
	}
	ip := c.RealIP()
	var received = []Post{
		{
			Name:    c.FormValue("guestName"),
			Message: c.FormValue("guestText"),
			Id:      ip,
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
	for i := 0; i < len(posts); i++ {
		if posts[i].Id == ip {
			return c.Render(http.StatusOK, "alert", "You can only post once! Delete your post to send a new one")
		}
	}
	result := append(received, posts...)
	write, _ := json.MarshalIndent(result, "", "	")
	os.WriteFile("internal/posts/posts.json", write, 0664)
	if err != nil {
		panic(err)
	}

	return c.Render(http.StatusOK, "guestbookpost", result)
}
func GuestbookDelete(c echo.Context) error {
	ip := c.RealIP()
	file, err := os.OpenFile("internal/posts/posts.json", os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var posts []Post
	var result []Post
	json.Unmarshal(data, &posts)
	fmt.Println(posts[0])
	for i := 0; i < len(posts); i++ {
		if posts[i].Id == string(ip) {
			fmt.Println("sexo")
			continue
		}
		fmt.Println(len(posts))
		var item = []Post{{
			Name:    posts[i].Name,
			Message: posts[i].Message,
			Id:      posts[i].Id,
		}}
		result = append(item, result...)
	}
	write, _ := json.MarshalIndent(result, "", "	")
	os.WriteFile("internal/posts/posts.json", write, 0664)
	if err != nil {
		panic(err)
	}
	return GetGuestbookentries(c)
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
