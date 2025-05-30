package home

import (
	"net/http"
	"shitty-portfolio/internal/handlers"
	"shitty-portfolio/internal/views"
)

type Handler struct{}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	return handlers.RenderWithDefaultLayout(w, r, views.Home())
}

func (h *Handler) RenderTemplate(w http.ResponseWriter, r *http.Request) error {
	return handlers.Render(w, r, views.Home())
}
