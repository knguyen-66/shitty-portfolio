package handlers

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"shitty-portfolio/internal/views"

	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, t templ.Component) error {
	return t.Render(r.Context(), w)
}

func RenderWithDefaultLayout(w http.ResponseWriter, r *http.Request, t templ.Component) error {
	prod, err := strconv.ParseBool(os.Getenv("APP_PROD"))
	if err != nil {
		slog.Error("app error", "err", err)
	}
	return views.DEFAULT_LAYOUT(t, prod).Render(r.Context(), w)
}
