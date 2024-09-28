package main

import (
	"fmt"
	"go-breeders/models"
	"go-breeders/pets"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// Get the breed
	breed, err := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Get the dog of the month from database
	dom, err := app.App.Models.Dog.GetDogOfMonthByID(1)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	// Create dog and decorate it
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			Name:             "Sam",
			BreedID:          breed.ID,
			Color:            "Black & Tan",
			DateOfBirth:      dob,
			SpayedOrNeutered: 0,
			Description:      "Sam is a very good boy.",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	// Serve the web page
	data := make(map[string]any)
	data["dog"] = dog

	app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// create a dog using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetMinWeight(10).
		SetMaxWeight(30).
		SetDescription("A mixed breed of unknow origin.").
		SetColor("Block and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// create a dog using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("felis slivestris catus").
		SetWeight(4).
		SetMinWeight(3).
		SetMaxWeight(8).
		SetDescription("A beautiful house cat.").
		SetColor("Calico").
		SetAge(1).
		SetGeographicOrigin("Canada").
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	catBreeds, err := app.App.CatService.GetAllCatBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// get species from URL itself.
	species := chi.URLParam(r, "species")

	// get breed from the URL.
	breed, err := url.QueryUnescape(chi.URLParam(r, "breed"))
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	fmt.Println("Species:", species, "Breed:", breed)

	// create a pet from abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// write the result as josn
	_ = t.WriteJSON(w, http.StatusOK, pet)
}
