package handlers

import (
	"context"
	"fmt"
	"net/http"
	"shitty-portfolio/data"
	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
)

func _handleBlogs() (templ.Component, error) {
	queries := data.New(data.DB)
	ctx := context.Background()
	blogs, err := queries.SelectBlogs(ctx)

	blogTagMap := make(map[int64][]data.BlogTag)
	for _, blog := range blogs {
		blogTags, err := queries.SelectBlogTagsByBlog(ctx, blog.ID)
		if err != nil {
			blogTagMap[blog.ID] = []data.BlogTag{}
			continue
		}
		blogTagMap[blog.ID] = blogTags
	}
	return views.Blogs(blogs, blogTagMap), err
}

func HandleBlogsPage(w http.ResponseWriter, r *http.Request) error {
	viewTemplate, err := _handleBlogs()
	if err != nil {
		fmt.Println("Error happened: " + err.Error())
	}
	return RenderWithDefaultLayout(w, r, viewTemplate)
}

func HandleBlogsTemplate(w http.ResponseWriter, r *http.Request) error {
	viewTemplate, err := _handleBlogs()
	if err != nil {
		fmt.Println("Error happened: " + err.Error())
	}
	return Render(w, r, viewTemplate)
}
