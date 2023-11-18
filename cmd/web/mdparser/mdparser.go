package mdparser

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/yuin/goldmark"
)

func ParseNewFiles() {
	file, _ := os.OpenFile("internal/markdown/This is a test.md", os.O_RDWR, 0664)
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	md := goldmark.New(goldmark.WithRendererOptions())
	var buf bytes.Buffer
	if err := md.Convert(data, &buf); err != nil {
		panic("failed to convert")
	}

	fmt.Println(string(data))

}
