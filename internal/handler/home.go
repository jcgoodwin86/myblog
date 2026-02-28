package handler

import (
	"net/http"

	"github.com/jcgoodwin/myblog/templates/pages"
)

func (app App) HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pages.Home().Render(r.Context(), w)
}
