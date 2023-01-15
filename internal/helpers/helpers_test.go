package helpers

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mardybom/covidsum/internal/models"
	"github.com/mardybom/covidsum/internal/utils"
)

func TestSummarizeCases(t *testing.T) {
	testData := []models.Case{
		{
			Age:      utils.NullableInt{true, false, 51},
			Province: "Phrae"},
		{
			Age:      utils.NullableInt{true, false, 21},
			Province: "Chumphon"},
		{
			Age:      utils.NullableInt{true, false, 62},
			Province: "Bangkok"},
		{
			Age:      utils.NullableInt{true, true, 0},
			Province: ""},
	}

	expectedData := models.Summary{
		Province: map[string]int{"Bangkok": 1, "Phrae": 2, "Chumphon": 1, "N/A": 1},
		AgeGroup: map[string]int{"0-30": 1, "31-60": 1, "60+": 1, "N/A": 1},
	}

	result := SummarizeCases(testData)

	eqProvince := reflect.DeepEqual(result.Province, expectedData.Province)
	eqAgeGroup := reflect.DeepEqual(result.AgeGroup, expectedData.AgeGroup)

	if !eqProvince || !eqAgeGroup {
		r, err := json.Marshal(result)
		if err != nil {
			t.Log(err)
		}
		e, err := json.Marshal(expectedData)
		if err != nil {
			t.Log(err)
		}
		t.Errorf("expected %s, but got %s", string(r), string(e))
	}

}
