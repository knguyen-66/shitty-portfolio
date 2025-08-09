package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"shitty-portfolio/data"
	"shitty-portfolio/internal/utils"
)

func main() {
	data.InitDatabaseDefault()
	defer data.CloseDatabase()

	entries, err := os.ReadDir("./data/blogs/")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fmt.Println("Setting up blog data for file: " + entry.Name())
		_, metadata, err := utils.ParseMarkdownBlog(entry.Name())
		if err != nil {
			fmt.Println("Parsing file error: " + err.Error())
			continue
		}
		fmt.Println(metadata.Title)

		ctx := context.Background()
		queries := data.New(data.DB)
		existedBlog, _ := queries.SelectBlogByContentFilename(ctx, entry.Name())
		if existedBlog.ID != 0 {
			fmt.Println("Blog existed. Skipping.")
			continue
		}
		newBlog := data.InsertBlogParams{
			Title:           metadata.Title,
			Slug:            utils.CreateSlug(metadata.Title),
			Summary:         sql.NullString{String: metadata.Summary, Valid: true},
			ThumbnailDir:    sql.NullString{String: "", Valid: true},
			ContentFilename: entry.Name(),
			TimeCreated:     sql.NullTime{Time: metadata.Date, Valid: true},
		}
		insertedBlog, err := queries.InsertBlog(ctx, newBlog)
		if err != nil {
			fmt.Println("Adding blog error: " + err.Error() + ". Skipping.")
			continue
		}
		if len(metadata.Warnings) > 0 {
			fmt.Println("Warnings:", metadata.Warnings)
		}
		fmt.Println("Inserted blog:", insertedBlog.ID, insertedBlog.Title, insertedBlog.Slug)

		for _, tagName := range metadata.Tags {
			newBlogTag := data.InsertBlogTagParams{
				IDBlog:  insertedBlog.ID,
				TagName: tagName,
			}
			insertedBlogTag, _ := queries.InsertBlogTag(ctx, newBlogTag)
			if err != nil {
				fmt.Println("Adding blog tag error: " + err.Error())
			}
			fmt.Println("Inserted blog tag:", insertedBlogTag.ID, insertedBlogTag.TagName, "for blog ID", insertedBlogTag.IDBlog)
		}
	}
}
