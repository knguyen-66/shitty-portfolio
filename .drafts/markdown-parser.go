package main

import (
	"bytes"
	// "context"
	// "database/sql"
	"fmt"
	"os"
	"reflect"
	"regexp"

	// "shitty-portfolio/data"
	// "shitty-portfolio/internal/utils"
	"time"

	figure "github.com/mangoumbrella/goldmark-figure"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var markdownConverter goldmark.Markdown
var buf bytes.Buffer
var parserContext parser.Context

func getKey(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		return fmt.Sprint(val)
	}
	return ""
}

func main() {
	source, err := os.ReadFile("./data/blogs/01-building-a-kinda-high-performance-chat-application-with-python-and-golang.md")
	if err != nil {
		panic(err)
	}
	markdownConverter = goldmark.New(goldmark.WithExtensions(extension.GFM, figure.Figure, meta.Meta))
	parserContext = parser.NewContext()
	if err := markdownConverter.Convert(source, &buf, parser.WithContext(parserContext)); err != nil {
		panic(err)
	}
	metadata := meta.Get(parserContext)
	re := regexp.MustCompile(`(?m)^<img`)
	blogStr := re.ReplaceAllString(buf.String(), "<img loading='lazy'")
	date, _ := time.Parse("02/01/2006", fmt.Sprint(metadata["Date"]))

	fmt.Println("Title:", metadata["Title"])
	fmt.Println("Summary:", metadata["Summary"])
	fmt.Println("Date:", date)
	fmt.Println("Tags:", metadata["Tags"], reflect.TypeOf(metadata["Tags"]))
	fmt.Println("----------")
	_ = blogStr
	// fmt.Println(blogStr)

	// db, err := sql.Open("sqlite3", "../app.db")
	// defer db.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// ctx := context.Background()
	// queries := data.New(db)
	// newBlog := data.InsertBlogParams{
	// 	Title:           fmt.Sprint(metadata["Title"]),
	// 	Slug:            utils.CreateSlug(fmt.Sprint(metadata["Title"])),
	// 	Summary:         sql.NullString{String: fmt.Sprint(metadata["Summary"]), Valid: true},
	// 	ThumbnailDir:    sql.NullString{String: "", Valid: true},
	// 	ContentFilename: "01-building-a-kinda-high-performance-chat-application-with-python-and-golang.md",
	// 	TimeCreated:     sql.NullTime{Time: date, Valid: true},
	// }
	// insertedBlog, err := queries.InsertBlog(ctx, newBlog)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Inserted blog:", insertedBlog.ID, insertedBlog.Title, insertedBlog.Slug)
	// for _, tagName := range metadata["Tags"].([]interface{}) {
	// 	newBlogTag := data.InsertBlogTagParams{
	// 		IDBlog:  insertedBlog.ID,
	// 		TagName: tagName.(string),
	// 	}
	// 	insertedBlogTag, err := queries.InsertBlogTag(ctx, newBlogTag)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	fmt.Println("Inserted blog tag:", insertedBlogTag.ID, insertedBlogTag.TagName, "for blog ID", insertedBlogTag.IDBlog)
	// }
}
