package blog_article

import (
	"context"
	"fmt"
	"net/http"
	"shitty-portfolio/data"
	"shitty-portfolio/internal/handler"
	"shitty-portfolio/internal/utils"
	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*handler.BaseHTTPHandler
	SlugName string
}

func (h *Handler) handle() (templ.Component, error) {
	queries := data.New(data.DB)
	ctx := context.Background()

	existedBlog, err := queries.SelectBlogBySlug(ctx, h.SlugName)
	if err != nil {
		return nil, err
	}
	if existedBlog.ID == 0 {
		return nil, fmt.Errorf("page handle: blog not found")
	}
	blogTags, err := queries.SelectBlogTagsByBlog(ctx, existedBlog.ID)
	if err != nil {
		return nil, fmt.Errorf("page handle: %w", err)
	}
	blogContent, _, err := utils.ParseMarkdownBlog(existedBlog.ContentFilename)
	if err != nil {
		return nil, fmt.Errorf("page handle: %w", err)
	}
	return views.BlogArticle(existedBlog, blogTags, blogContent), nil
}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	h.SlugName = chi.URLParam(r, "articleSlug")
	viewTemplate, err := h.handle()
	if err != nil {
		return err
	}
	return handler.RenderWithDefaultLayout(w, r, viewTemplate)
}

func (h *Handler) RenderTemplate(w http.ResponseWriter, r *http.Request) error {
	h.SlugName = chi.URLParam(r, "articleSlug")
	viewTemplate, err := h.handle()
	if err != nil {
		return err
	}
	return handler.Render(w, r, viewTemplate)
}
