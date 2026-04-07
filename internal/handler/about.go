package handler

import (
	"net/http"

	"github.com/jcgoodwin/myblog/templates/pages"
)

func (app App) HandleAbout(w http.ResponseWriter, r *http.Request) {
	if err := pages.About().Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
