package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jcgoodwin/myblog/internal/handler"
)

func main() {
	port := "8080"
	app := handler.App{}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.HandleHome)
	mux.HandleFunc("GET /posts/{slug}", app.HandlePost)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
