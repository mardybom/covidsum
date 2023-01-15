package repository

import "github.com/mardybom/covidsum/internal/models"

type ApiRepo interface {
	GetCases() ([]models.Case, error)
}
