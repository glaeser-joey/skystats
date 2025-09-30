package data

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"strings"
	"sync"
)

//go:embed airlines.csv
var airlinesCSV []byte

type Airline struct {
	Code string
	Name string
	ICAO string
	IATA string
}

var (
	airlinesOnce  sync.Once
	airlinesIndex map[string]Airline
)

func loadAirlines() {
	airlinesIndex = make(map[string]Airline)
	reader := csv.NewReader(bytes.NewReader(airlinesCSV))

	records, err := reader.ReadAll()
	if err != nil || len(records) < 2 {
		return
	}

	records[0][0] = strings.TrimPrefix(records[0][0], "\ufeff")

	cols := make(map[string]int)
	for i, name := range records[0] {
		cols[strings.ToLower(strings.TrimSpace(name))] = i
	}

	for _, row := range records[1:] {
		airline := Airline{
			Code: getValue(row, cols["code"]),
			Name: getValue(row, cols["name"]),
			ICAO: getValue(row, cols["icao"]),
			IATA: getValue(row, cols["iata"]),
		}

		if airline.Name == "" {
			continue
		}

		for _, key := range []string{airline.Code, airline.ICAO, airline.IATA} {
			if key = strings.ToUpper(strings.TrimSpace(key)); key != "" {
				if _, exists := airlinesIndex[key]; !exists {
					airlinesIndex[key] = airline
				}
			}
		}
	}
}

func getValue(row []string, idx int) string {
	if idx >= 0 && idx < len(row) {
		return strings.TrimSpace(row[idx])
	}
	return ""
}

func LookupAirline(code string) (Airline, bool) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if code == "" {
		return Airline{}, false
	}

	airlinesOnce.Do(loadAirlines)
	airline, ok := airlinesIndex[code]
	return airline, ok
}
