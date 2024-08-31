package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"shitty-portfolio/internal/handlers"
)

func LoadRouterPaths(r *chi.Mux) {
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./internal/static"))))

	// full pages
	r.Get("/", handlers.Make(handlers.HandleHomePage))
	r.Get("/blogs", handlers.Make(handlers.HandleBlogsPage))
	r.Get("/work-archive", handlers.Make(handlers.HandleWorkArchivePage))

	// page templates, for HTMX
	r.Route("/templates", func(r chi.Router) {
		r.Get("/home", handlers.Make(handlers.HandleHomeTemplate))
		r.Get("/blogs", handlers.Make(handlers.HandleBlogsTemplate))
		r.Get("/work-archive", handlers.Make(handlers.HandleWorkArchiveTemplate))
	})
}
