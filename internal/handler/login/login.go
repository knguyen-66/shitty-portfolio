package login

import (
	"fmt"
	"net/http"
	"shitty-portfolio/internal/handler"
	"shitty-portfolio/internal/views"
)

type Handler struct {
	*handler.BaseHTTPHandler
}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		// handle login
		if err := r.ParseForm(); err != nil {
			handler.RenderWithDefaultLayout(w, r, views.Login("Error parsing the form"))
			return fmt.Errorf("login: error parsing form: %w", err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		if len(username) < 8 || len(password) < 8 {
			handler.RenderWithDefaultLayout(w, r, views.Login("Why so short?"))
			return nil
		}

		// if r.FormValue("password")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	handler.RenderWithDefaultLayout(w, r, views.Login(""))
	return nil
}
