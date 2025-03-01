package blogs

import (
	"context"
	"fmt"
	"net/http"
	"shitty-portfolio/data"
	"shitty-portfolio/internal/handlers"
	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
)

type Handler struct{}

func handle() (templ.Component, error) {
	queries := data.New(data.DB)
	ctx := context.Background()
	blogs, err := queries.SelectBlogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("page handle: %w", err)
	}
	blogTagMap := make(map[int64][]data.BlogTag)
	for _, blog := range blogs {
		blogTags, err := queries.SelectBlogTagsByBlog(ctx, blog.ID)
		if err != nil {
			return nil, fmt.Errorf("page handle: %w", err)
		}
		blogTagMap[blog.ID] = blogTags
	}
	return views.Blogs(blogs, blogTagMap), err
}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	viewTemplate, err := handle()
	if err != nil {
		return err
	}
	return handlers.RenderWithDefaultLayout(w, r, viewTemplate)
}

func (h *Handler) RenderTemplate(w http.ResponseWriter, r *http.Request) error {
	viewTemplate, err := handle()
	if err != nil {
		return err
	}
	return handlers.Render(w, r, viewTemplate)
}
