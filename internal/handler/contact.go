package handler

import (
	"net/http"

	"github.com/jcgoodwin/myblog/templates/pages"
)

func (app App) HandleContact(w http.ResponseWriter, r *http.Request) {
	if err := pages.Contact().Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
