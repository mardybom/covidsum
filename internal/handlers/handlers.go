package handlers

import (
	"github.com/mardybom/covidsum/internal/config"
	"github.com/mardybom/covidsum/internal/repository"
	"github.com/mardybom/covidsum/internal/repository/apirepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	Api repository.ApiRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		Api: apirepo.NewApiRepo(a),
	}
}
