package config

import (
	"database/sql"
	adadapters "go-breeders/adapters"
	"go-breeders/models"
	"sync"
)

type Application struct {
	Models     *models.Models
	CatService *adadapters.RemoteService
}

var instance *Application
var once sync.Once
var db *sql.DB
var catService *adadapters.RemoteService

func New(pool *sql.DB, cs *adadapters.RemoteService) *Application {
	db = pool
	catService = cs
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models:     models.New(db),
			CatService: catService,
		}
	})
	return instance
}
