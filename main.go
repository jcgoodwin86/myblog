package main

import (
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/jcgoodwin/myblog/internal/handler"
	"github.com/jcgoodwin/myblog/internal/middleware"
	"github.com/jcgoodwin/myblog/internal/model"
)

func main() {
	// Prerender all the posts pages
	if err := model.LoadAllPosts(); err != nil {
		log.Fatal(err)
	}

	mime.AddExtensionType(".css", "text/css; charset=utf-8")
	mime.AddExtensionType(".js", "application/javascript")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := handler.App{}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.HandleHome)
	mux.HandleFunc("GET /about", app.HandleAbout)
	mux.HandleFunc("GET /contact", app.HandleContact)
	mux.HandleFunc("GET /posts/{slug}", app.HandlePost)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      middleware.SetContentType(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
