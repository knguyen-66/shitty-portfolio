package handlers

import (
	"log/slog"
	"net/http"

	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
)

type HTTPHandler interface {
	RenderPage(w http.ResponseWriter, r *http.Request) error
	RenderTemplate(w http.ResponseWriter, r *http.Request) error
}

func Make(h HTTPHandler, isTemplate bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if isTemplate {
			err = h.RenderTemplate(w, r)
		} else {
			err = h.RenderPage(w, r)
		}
		if err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, t templ.Component) error {
	return t.Render(r.Context(), w)
}

func RenderWithDefaultLayout(w http.ResponseWriter, r *http.Request, t templ.Component) error {
	return views.DEFAULT_LAYOUT(t).Render(r.Context(), w)
}
