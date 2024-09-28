package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
)

func main() {
	source, err := os.ReadFile("../internal/data/blogs/01-building-a-kinda-high-performance-chat-application-with-python-and-golang.md")
	if err != nil {
		panic(err)
	}
	markdownConverter := goldmark.New(goldmark.WithExtensions(extension.GFM, figure.Figure, meta.Meta))
	var buf bytes.Buffer
	if err := markdownConverter.Convert(source, &buf); err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`<img(.*)>`)
	blogStr := re.ReplaceAllString(buf.String(), "<img loading='lazy'$1>")
	fmt.Println(blogStr)
}
