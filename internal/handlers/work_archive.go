package handlers

import (
	"net/http"
	"shitty-portfolio/internal/views"
)

func HandleWorkArchivePage(w http.ResponseWriter, r *http.Request) error {
	return RenderWithDefaultLayout(w, r, views.WorkArchive())
}

func HandleWorkArchiveTemplate(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.WorkArchive())
}
