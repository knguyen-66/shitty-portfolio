package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
)

type HTTPHandler interface {
	IsProtected() bool

	RenderPage(w http.ResponseWriter, r *http.Request) error
	RenderTemplate(w http.ResponseWriter, r *http.Request) error
}

func Make(h HTTPHandler, isTemplate bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if h.IsProtected() {
			// for now
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		if isTemplate {
			err = h.RenderTemplate(w, r)
		} else {
			err = h.RenderPage(w, r)
		}
		if err != nil {
			// output error with type of handler
			if os.Getenv("APP_PRODUCTION") == "false" {
				RenderWithDefaultLayout(w, r, views.NotFound(err.Error()))
			}
			slog.Error("HTTP handler error", "handler", fmt.Sprintf("%T", h), "err", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, t templ.Component) error {
	return t.Render(r.Context(), w)
}

func RenderWithDefaultLayout(w http.ResponseWriter, r *http.Request, t templ.Component) error {
	return views.DEFAULT_LAYOUT(t).Render(r.Context(), w)
}

// default implementation of HTTPHandler

type BaseHTTPHandler struct{}

func (h *BaseHTTPHandler) IsProtected() bool {
	return false
}

func (h *BaseHTTPHandler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("<h1>Page not implemented</h1>"))
	return fmt.Errorf("not implemented")
}

func (h *BaseHTTPHandler) RenderTemplate(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("<h1>Page not implemented</h1>"))
	return fmt.Errorf("not implemented")
}
