package utils

import (
	"bytes"
	"errors"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"time"

	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var ALLOWED_BLOG_METADATA = []string{"Title", "Summary", "Date", "Tags"}

type BlogMetadata struct {
	Title    string
	Summary  string
	Date     time.Time
	Tags     []string
	Warnings []string
}

var markdownConverter goldmark.Markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM, figure.Figure.WithImageLink(), meta.Meta),
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

func ParseMarkdownBlog(filename string) (string, *BlogMetadata, error) {
	source, err := os.ReadFile("data/blogs/" + filename)
	if err != nil {
		return "", nil, err
	}
	var buf bytes.Buffer
	parserContext := parser.NewContext()
	if err := markdownConverter.Convert(source, &buf, parser.WithContext(parserContext)); err != nil {
		return "", nil, err
	}
	markdownMetadata := meta.Get(parserContext)
	blogMetadata := BlogMetadata{}

	title, exists := markdownMetadata["Title"].(string)
	if !exists {
		return "", nil, errors.New("markdown parse error: Title metadata is required")
	}
	summary, exists := markdownMetadata["Summary"].(string)
	if !exists {
		blogMetadata.Warnings = append(blogMetadata.Warnings, "Summary metadata is empty")
	}
	dateStr, exists := markdownMetadata["Date"].(string)
	if !exists {
		blogMetadata.Warnings = append(blogMetadata.Warnings, "Date metadata is empty")
	}
	date, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		blogMetadata.Warnings = append(blogMetadata.Warnings, "Date metadata must be in the format of dd/mm/yyyy")
	}
	tags, exists := markdownMetadata["Tags"].([]interface{})
	if !exists {
		blogMetadata.Warnings = append(blogMetadata.Warnings, "Tags metadata is empty")
	}

	blogMetadata.Title = title
	blogMetadata.Summary = summary
	blogMetadata.Date = date
	for _, tag := range tags {
		blogMetadata.Tags = append(blogMetadata.Tags, tag.(string))
	}

	re := regexp.MustCompile(`(?m)^<img`)
	blogContent := re.ReplaceAllString(buf.String(), "<img loading='lazy'")
	return blogContent, &blogMetadata, nil
}
