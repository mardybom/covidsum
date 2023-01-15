package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mardybom/covidsum/internal/config"
	"github.com/mardybom/covidsum/internal/handlers"
	"github.com/mardybom/covidsum/internal/helpers"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	run()
}

func run() {
	// read flag
	dataSrcUrl := flag.String("datasrcurl", "http://static.wongnai.com/devinterview/covid-cases.json", "Datasource URL")

	app.DataSourceUrl = *dataSrcUrl

	// setup logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := handlers.NewRepo(&app)
	helpers.NewHelpers(&app)

	r := gin.Default()
	r.GET("/covid/cases", func(c *gin.Context) {
		cases, err := repo.Api.GetCases()
		if err != nil {
			app.ErrorLog.Fatal(err)
		}

		c.JSON(200, gin.H{
			"Data": cases,
		})
	})
	r.GET("/covid/summary", func(c *gin.Context) {
		cases, err := repo.Api.GetCases()
		if err != nil {
			app.ErrorLog.Fatal(err)
		}

		sum := helpers.SummarizeCases(cases)

		c.JSON(200, gin.H{
			"Data": sum,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
