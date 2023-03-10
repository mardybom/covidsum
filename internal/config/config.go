package config

import (
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	DataSourceUrl string
}
