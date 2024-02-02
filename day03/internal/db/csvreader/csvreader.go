package csvreader

import (
	"day03/internal/types"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func LoadDataFromCSV(filePath string) []types.Place {
	var places []types.Place
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %s", err)
		return nil
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %s", err)
		return nil
	}

	for i, record := range records[1:] {
		lon, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Printf("Error parsing longitude: %s", err)
			continue
		}

		lat, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Printf("Error parsing latitude: %s", err)
			continue
		}

		places = append(places, types.Place{
			ID:      i + 1,
			Name:    record[1],
			Address: record[2],
			Phone:   record[3],
			Location: struct {
				Lon float64 `json:"lon"`
				Lat float64 `json:"lat"`
			}{Lon: lon, Lat: lat},
		})
	}
	return places
}
