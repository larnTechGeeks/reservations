package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/larnTechGeeks/reservations/pkg/config"
	"github.com/larnTechGeeks/reservations/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// routes
	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)

	// build a File Server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
