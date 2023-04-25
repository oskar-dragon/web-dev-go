package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"router/views"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tmpl, _ := views.Parse(filepath)
	tmpl.Execute(w, nil)
}
func homeHander(w http.ResponseWriter, r* http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r* http.Request){
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func notFoundhandler(w http.ResponseWriter, r* http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1>")
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	return r
}

func articleHandler(w http.ResponseWriter, r* http.Request) {
	articleId := chi.URLParam(r, "articleId")

	fmt.Fprintf(w, "<h1>Article number" + articleId + "</h1>")
}

func main() {
	r := setupRouter()

	r.Get("/", homeHander)
	r.Get("/contact", contactHandler)
	r.NotFound(notFoundhandler)

	r.Get("/articles/{articleId}", articleHandler)

	http.ListenAndServe(":3000", r)
}