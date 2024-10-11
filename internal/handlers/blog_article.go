package handlers

import (
	"context"
	"fmt"
	"net/http"
	"shitty-portfolio/data"
	"shitty-portfolio/internal/utils"
	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func _handleBlogArticle(slugName string) (templ.Component, error) {
	queries := data.New(data.DB)
	ctx := context.Background()

	existedBlog, err := queries.SelectBlogBySlug(ctx, slugName)
	if err != nil {
		return nil, err
	}
	if existedBlog.ID == 0 {
		return nil, fmt.Errorf("blog not found")
	}
	blogTags, err := queries.SelectBlogTagsByBlog(ctx, existedBlog.ID)
	if err != nil {
		return nil, err
	}
	blogContent, _, err := utils.ParseMarkdownBlogWithMetadata(existedBlog.ContentFilename)
	if err != nil {
		return nil, err
	}
	return views.BlogArticle(existedBlog, blogTags, blogContent), nil
} 

func HandleBlogArticlePage(w http.ResponseWriter, r *http.Request) error {
	slugName := chi.URLParam(r, "articleSlug")
	viewTemplate, err := _handleBlogArticle(slugName)
	if err != nil {
		fmt.Println("Error happened: " + err.Error())
	}
	return RenderWithDefaultLayout(w, r, viewTemplate)
}

func HandleBlogArticleTemplate(w http.ResponseWriter, r *http.Request) error {
	slugName := chi.URLParam(r, "articleSlug")
	viewTemplate, err := _handleBlogArticle(slugName)
	if err != nil {
		fmt.Println("Error happened: " + err.Error())
	}
	return Render(w, r, viewTemplate)
}
