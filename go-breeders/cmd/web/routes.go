package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", app.ShowHome)
	mux.Get("/dog-of-month", app.DogOfMonth)

	mux.Get("/{page}", app.ShowPage)

	// factory routes
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)
	mux.Get("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
	mux.Get("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)

	// builder routes
	mux.Get("/api/dog-from-builder", app.CreateDogWithBuilder)
	mux.Get("/api/cat-from-builder", app.CreateCatWithBuilder)

	// adapter  routes
	mux.Get("/api/cat-breeds", app.GetAllCatBreeds)

	// get data from database
	mux.Get("/api/dog-breeds", app.GetAllDogBreedsJSON)

	mux.Get("/api/animal-from-abstract-factory/{species}/{breed}", app.AnimalFromAbstractFactory)

	return mux
}
