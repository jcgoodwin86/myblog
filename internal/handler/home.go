package handler

import (
	"net/http"

	"github.com/jcgoodwin/myblog/internal/model"
	"github.com/jcgoodwin/myblog/templates/pages"
)

func (app App) HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	posts := model.GetAllPosts()
	canonical := "https://oddcodeout.io" + r.URL.Path

	if err := pages.Home(posts, canonical).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
