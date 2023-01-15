package apirepo

import (
	"github.com/mardybom/covidsum/internal/config"
	"github.com/mardybom/covidsum/internal/repository"
)

type apiRepo struct {
	App           *config.AppConfig
	DataSourceUrl string
}

func NewApiRepo(a *config.AppConfig) repository.ApiRepo {
	return &apiRepo{
		App:           a,
		DataSourceUrl: a.DataSourceUrl,
	}
}
