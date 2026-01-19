package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	repository "github.com/jeevaprakashdr/image-gallery/infrastructure/postgres/sqlc"
	"github.com/jeevaprakashdr/image-gallery/services/images"
)

// Mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good for now!!"))
	})

	r.Route("/images", func(r chi.Router) {
		imageService := images.NewService(repository.New(app.db))
		imageHandler := images.NewHandler(imageService, app.wsClientConnection)

		r.Get("/", imageHandler.ListImages)
		r.Post("/upload", imageHandler.SaveImage)
		r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
			tag := r.URL.Query().Get("tag")
			imageHandler.SearchImages(tag, w, r)
		})
	})

	return r
}

type application struct {
	config             config
	db                 *pgx.Conn
	wsClientConnection *websocket.Conn
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
