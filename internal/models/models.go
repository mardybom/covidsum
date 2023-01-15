package models

import "github.com/mardybom/covidsum/internal/utils"

type Case struct {
	ConfirmDate    string
	No             string
	Age            utils.NullableInt
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     int
	District       string
	ProvinceEn     string
	StatQuarantine int
}

type Data struct {
	Data []Case
}

type Summary struct {
	Province map[string]int
	AgeGroup map[string]int
}
