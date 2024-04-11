package SpecReader

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Blips struct {
	Blips []Blip
}

type Blip struct {
	Name        string `csv:"name"`
	Quadrant    string `csv:"quadrant"`
	Ring        string `csv:"ring"`
	IsNew       bool   `csv:"isNew"`
	Moved       int8   `csv:"moved"`
	Description string `csv:"description"`
}

// Read csv spec file and create Blips from that
// Code inspired from:
// https://stackoverflow.com/questions/20768511/unmarshal-csv-record-into-struct-in-go
func ReadCsvSpec(filePath string) Blips {
	// Open file
	csvFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Opened csv file!")
	defer csvFile.Close()

	// Read file
	var smallBlips []Blip
	err = gocsv.UnmarshalFile(csvFile, &smallBlips)
	if err != nil {
		panic(err)
	}

	blips := Blips{
		Blips: smallBlips,
	}

	return blips
}

func CsvToString(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// Build the CSV string
	csvString := string(data)

	return csvString
}
