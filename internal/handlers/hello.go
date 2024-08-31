package handlers

import (
	"net/http"
	"shitty-portfolio/internal/views"
)

func HandleHello(w http.ResponseWriter, r *http.Request) error {
	// return views.Hello().Render(r.Context(), w)
	return Render(w, r, views.Hello())
}
