package handler

import (
	"log"
	"net/http"
	"os"
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
		log.Println(err)
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	if err := pages.Post(*postData).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
