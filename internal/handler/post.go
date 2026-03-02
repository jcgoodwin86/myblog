package handler

import (
	"net/http"
	"strings"

	"github.com/jcgoodwin/myblog/internal/model"
	"github.com/jcgoodwin/myblog/templates/pages"
)

func (app App) HandlePost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	// sanitize the slug
	if strings.ContainsAny(slug, "./\\") {
		http.NotFound(w, r)
		return
	}

	postData, err := model.LoadPost(slug)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pages.Post(*postData).Render(r.Context(), w)
}
