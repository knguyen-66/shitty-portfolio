package admin

import (
	"net/http"
	"shitty-portfolio/internal/handler"
)

type Handler struct {
	*handler.BaseHTTPHandler
}

func (h *Handler) IsProtected() bool {
	return true
}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("<h1>Admin page</h1>"))
	return nil
}

func (h *Handler) RenderTemplate(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("<h1>Admin template</h1>"))
	return nil
}
