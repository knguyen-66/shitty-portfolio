package handlers

import (
	"net/http"
	"shitty-portfolio/internal/views"
)

func HandleBlogsPage(w http.ResponseWriter, r *http.Request) error {
	return RenderWithDefaultLayout(w, r, views.Blogs())
}

func HandleBlogsTemplate(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Blogs())
}
