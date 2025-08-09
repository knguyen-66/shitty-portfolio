package work_archive

import (
	"net/http"
	"shitty-portfolio/internal/handler"
	"shitty-portfolio/internal/views"
)

type Handler struct {
	*handler.BaseHTTPHandler
}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	return handler.RenderWithDefaultLayout(w, r, views.WorkArchive())
}

func (h *Handler) RenderTemplate(w http.ResponseWriter, r *http.Request) error {
	return handler.Render(w, r, views.WorkArchive())
}
