package routes

import (
	"net/http"

	"shitty-portfolio/internal/handler"
	"shitty-portfolio/internal/handler/admin"
	"shitty-portfolio/internal/handler/blog_article"
	"shitty-portfolio/internal/handler/blogs"
	"shitty-portfolio/internal/handler/home"
	"shitty-portfolio/internal/handler/login"
	"shitty-portfolio/internal/handler/work_archive"
	"shitty-portfolio/internal/views"

	"github.com/go-chi/chi/v5"
)

func LoadRouterPaths(r *chi.Mux) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		handler.RenderWithDefaultLayout(w, r, views.NotFound(""))
	})

	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/static/imgs/favicon.ico")
	})
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./internal/static"))))

	// full pages
	r.HandleFunc("/login", handler.Make(&login.Handler{}, false))
	r.Get("/admin", handler.Make(&admin.Handler{}, false))

	r.Get("/", handler.Make(&home.Handler{}, false))
	r.Get("/blogs", handler.Make(&blogs.Handler{}, false))
	r.Get("/blogs/{articleSlug:[a-z-]+}", handler.Make(&blog_article.Handler{}, false))
	r.Get("/work-archive", handler.Make(&work_archive.Handler{}, false))

	// page templates for HTMX
	r.Route("/templates", func(r chi.Router) {
		r.Get("/home", handler.Make(&home.Handler{}, true))
		r.Get("/blogs", handler.Make(&blogs.Handler{}, true))
		r.Get("/blogs/{articleSlug:[a-z-]+}", handler.Make(&blog_article.Handler{}, true))
		r.Get("/work-archive", handler.Make(&work_archive.Handler{}, true))
	})
}
