package handler

import (
	"net/http"

	"github.com/jcgoodwin/myblog/templates/pages"
)

func (app App) HandleAbout(w http.ResponseWriter, r *http.Request) {
	canonical := "https://oddcodeout.io" + r.URL.Path
	if err := pages.About(canonical).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
