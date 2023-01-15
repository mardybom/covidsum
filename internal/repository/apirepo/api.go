package apirepo

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mardybom/covidsum/internal/models"
)

func (m *apiRepo) GetCases() ([]models.Case, error) {
	var cases []models.Case

	resp, err := http.Get(m.App.DataSourceUrl)
	if err != nil {
		m.App.ErrorLog.Fatal(err)
		return cases, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		m.App.ErrorLog.Fatal(err)
		return cases, err
	}

	var data models.Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		m.App.ErrorLog.Fatal(err)
		return cases, err
	}
	cases = data.Data

	return cases, nil
}
