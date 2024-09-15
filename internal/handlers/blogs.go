package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"shitty-portfolio/internal/views"

	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	// "github.com/yuin/goldmark/renderer/html"
)

var markdownConverter goldmark.Markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM, figure.Figure, meta.Meta),
	// goldmark.WithRendererOptions(html.WithUnsafe()),
)

func renderBlogPost(filename string) (string, error) {
	source, err := os.ReadFile("internal/data/blogs/" + filename)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := markdownConverter.Convert(source, &buf); err != nil {
		return "", err
	}
	re := regexp.MustCompile(`<img(.*)>`)
	blogStr := re.ReplaceAllString(buf.String(), "<img loading='lazy'$1>")
	return blogStr, nil
}

func HandleBlogsPage(w http.ResponseWriter, r *http.Request) error {
	testPost, err := renderBlogPost("01-building-a-kinda-high-performance-chat-application-with-python-and-golang.md")
	if err != nil {
		fmt.Println("Error happened: " + err.Error())
	}
	return RenderWithDefaultLayout(w, r, views.Blogs(testPost))
}

func HandleBlogsTemplate(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Blogs(""))
}
