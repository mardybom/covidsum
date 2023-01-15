package helpers

import (
	"github.com/mardybom/covidsum/internal/config"
	"github.com/mardybom/covidsum/internal/models"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func SummarizeCases(cases []models.Case) models.Summary {
	var summary models.Summary
	province := make(map[string]int)
	ageGroup := make(map[string]int)

	for _, c := range cases {
		// categorize cases by province
		if c.Province != "" {
			province[c.Province] += 1
		} else {
			province["N/A"] += 1
		}

		// categorize cases by age
		if c.Age.Null {
			ageGroup["N/A"] += 1
		} else {
			switch {
			case c.Age.Value <= 30:
				ageGroup["0-30"] += 1

			case c.Age.Value >= 31 && c.Age.Value <= 60:
				ageGroup["31-60"] += 1

			case c.Age.Value >= 60:
				ageGroup["60+"] += 1
			default:
				ageGroup["N/A"] += 1
			}
		}
	}

	summary.Province = province
	summary.AgeGroup = ageGroup

	return summary
}
