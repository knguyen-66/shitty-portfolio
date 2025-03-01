package routes

import (
	"net/http"

	"shitty-portfolio/internal/handlers"
	"shitty-portfolio/internal/handlers/blog_article"
	"shitty-portfolio/internal/handlers/blogs"
	"shitty-portfolio/internal/handlers/home"
	"shitty-portfolio/internal/handlers/work_archive"

	"github.com/go-chi/chi/v5"
)

func LoadRouterPaths(r *chi.Mux) {
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/static/imgs/favicon.ico")
	})
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./internal/static"))))

	// full pages
	r.Get("/", handlers.Make(&home.Handler{}, false))
	r.Get("/blogs", handlers.Make(&blogs.Handler{}, false))
	r.Get("/blogs/{articleSlug:[a-z-]+}", handlers.Make(&blog_article.Handler{}, false))
	r.Get("/work-archive", handlers.Make(&work_archive.Handler{}, false))

	// page templates for HTMX
	r.Route("/templates", func(r chi.Router) {
		r.Get("/home", handlers.Make(&home.Handler{}, true))
		r.Get("/blogs", handlers.Make(&blogs.Handler{}, true))
		r.Get("/blogs/{articleSlug:[a-z-]+}", handlers.Make(&blog_article.Handler{}, true))
		r.Get("/work-archive", handlers.Make(&work_archive.Handler{}, true))
	})
}
