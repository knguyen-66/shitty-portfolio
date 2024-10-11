package utils

import (
	"bytes"
	"log/slog"
	"os"
	"regexp"
	"strings"

	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var markdownConverter goldmark.Markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM, figure.Figure, meta.Meta),
	// goldmark.WithRendererOptions(html.WithUnsafe()),
)

// https://gabrieleromanato.name/go-create-a-slug-from-a-string
func CreateSlug(input string) string {
	// Remove special characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		slog.Error("Error compiling regex: %v", "err", err.Error())
	}
	processedString := reg.ReplaceAllString(input, " ")

	// Remove leading and trailing spaces
	processedString = strings.TrimSpace(processedString)

	// Replace spaces with dashes
	slug := strings.ReplaceAll(processedString, " ", "-")

	// Convert to lowercase
	slug = strings.ToLower(slug)

	return slug
}

func ParseMarkdownBlogWithMetadata(filename string) (string, map[string]interface{}, error) {
	source, err := os.ReadFile("data/blogs/" + filename)
	if err != nil {
		return "", nil, err
	}
	var buf bytes.Buffer
	parserContext := parser.NewContext()
	if err := markdownConverter.Convert(source, &buf, parser.WithContext(parserContext)); err != nil {
		return "", nil, err
	}
	metadata := meta.Get(parserContext)
	// if metadata == nil {
	// 	return "", nil, nil
	// }
	re := regexp.MustCompile(`(?m)^<img`)
	blogContent := re.ReplaceAllString(buf.String(), "<img loading='lazy'")
	return blogContent, metadata, nil
}
