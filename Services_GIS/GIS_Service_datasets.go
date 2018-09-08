// Contains information for specific datasets

package Services_GIS

import (
	"encoding/csv"
	"os"
)

//
// city datset
//

// Specific to city dataset
// https://simplemaps.com/data/world-cities
var cityIndex = map[string]int{
	"name":       1,
	"longitude":  2,
	"latitude":   3,
	"population": 4,
	"country":    5,
}

func (s *GISService) loadCityDataset() {
	// Load basic data set (cities)
	csvFile, err := os.Open("Data\\cities.csv")
	checkError(err)
	cityData, err := csv.NewReader(csvFile).ReadAll()
	checkError(err)
	s.datasets = map[string]*dataset{}
	s.datasets["city"] = &dataset{name: "city", indices: cityIndex, data: cityData}
}

//
// country datset
//

var countryIndex = map[string]int{
	"name":       1,
	"population": 61,
}

func (s *GISService) loadCountryDataset() {
	// Load basic data set (cities)
	csvFile, err := os.Open("Data\\countries.csv")
	checkError(err)
	countryData, err := csv.NewReader(csvFile).ReadAll()
	checkError(err)
	s.datasets = map[string]*dataset{}
	s.datasets["country"] = &dataset{name: "country", indices: countryIndex, data: countryData}
}
