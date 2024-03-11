package SpecReader

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type Blips struct {
	Blips []Blip 
}

type Blip struct {
	Name     string `csv:"name"`
	Quadrant int8   `csv:"quadrant"`
	Ring     int8   `csv:"ring"`
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
