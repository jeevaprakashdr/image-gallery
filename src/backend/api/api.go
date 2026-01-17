package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jeevaprakashdr/image-gallery/internal/images"
	repository "github.com/jeevaprakashdr/image-gallery/postgres/sqlc"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good for now!!"))
	})

	imageService := images.NewService(repository.New(app.db))
	imageHandler := images.NewHandler(imageService)

	r.Route("/images", func(r chi.Router) {
		r.Get("/", imageHandler.ListImages)
		r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
			tag := r.URL.Query().Get("tag")
			imageHandler.SearchImages(tag, w, r)
		})
	})

	return r
}

type application struct {
	config config
	// logger
	db *pgx.Conn
}

// Run
func (app *application) run(h http.Handler) error {
	server := http.Server{
		Addr:    app.config.address,
		Handler: h,
	}

	log.Printf("server started at address %s", app.config.address)

	return server.ListenAndServe()
}

type config struct {
	address string
	db      dbConfig
}

type dbConfig struct {
	connectionString string
}
