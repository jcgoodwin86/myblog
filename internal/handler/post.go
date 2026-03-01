package handler

import (
	"net/http"
	"strings"

	"github.com/jcgoodwin/myblog/internal/model"
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Error"))
		return
	}

	w.Write([]byte(postData.Content))
}
