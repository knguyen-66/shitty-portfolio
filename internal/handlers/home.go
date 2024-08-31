package handlers

import (
	"net/http"
	"shitty-portfolio/internal/views"
)

func HandleHomePage(w http.ResponseWriter, r *http.Request) error {
	return RenderWithDefaultLayout(w, r, views.Home())
}

func HandleHomeTemplate(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Home())
}
