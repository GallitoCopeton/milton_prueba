package main

import (
	"challenge/routes/products"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Set up static fie serving
	staticFileServer := http.FileServer(http.Dir("static"))
	r.Handle("/*", http.StripPrefix("/", staticFileServer))
	r.Get("/", ServeIndex)
	r.Get("/index.html", ServeIndex)

	// Set up API v1
	r.Mount("/api/v1", products.Routes())

	// Spin up server
	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
