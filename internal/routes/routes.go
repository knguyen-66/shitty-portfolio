package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"shitty-portfolio/internal/handlers"
)

func handleFavicon(w http.ResponseWriter, r *http.Request) error {
	http.ServeFile(w, r, "./internal/static/imgs/favicon.ico")
	return nil
}

func LoadRouterPaths(r *chi.Mux) {
	r.Handle("/favicon.ico", handlers.Make(handleFavicon))
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./internal/static"))))

	// full pages
	r.Get("/", handlers.Make(handlers.HandleHomePage))
	r.Get("/blogs", handlers.Make(handlers.HandleBlogsPage))
	r.Get("/blogs/{articleSlug:[a-z-]+}", handlers.Make(handlers.HandleBlogArticlePage))
	r.Get("/work-archive", handlers.Make(handlers.HandleWorkArchivePage))

	// page templates for HTMX
	r.Route("/templates", func(r chi.Router) {
		r.Get("/home", handlers.Make(handlers.HandleHomeTemplate))
		r.Get("/blogs", handlers.Make(handlers.HandleBlogsTemplate))
		r.Get("/blogs/{articleSlug:[a-z-]+}", handlers.Make(handlers.HandleBlogArticleTemplate))
		r.Get("/work-archive", handlers.Make(handlers.HandleWorkArchiveTemplate))
	})
}
