package handler

import (
	"encoding/csv"
	"fmt"
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
	file, err := os.OpenFile("./internal/posts/posts.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	var latest error
	if err != nil {
		latest = err
		panic(err)
	}
	err = csv.NewWriter(file).WriteAll(post)
	if err != nil {
		latest = err
		panic(err)
	}
	reader := csv.NewReader(file)
	prePosts, err := reader.Read()
	fmt.Println(err)
	fmt.Println(prePosts)
	if err != nil {
		println("fodeuuu")
		panic(err)
	}
	fmt.Println("doneee")
	fmt.Println(latest)
	defer file.Close()
	return c.Render(http.StatusOK, "guestbookpost", name)
}
