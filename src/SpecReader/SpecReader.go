package SpecReader

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

// Code inspired from:
// https://tutorialedge.net/golang/parsing-json-with-golang/
type Blips struct {
	Blips []Blip `json:"Blips"`
}

type Blip struct {
	Name     string `json:"name" csv:"name"`
	Quadrant int8   `json:"quadrant" csv:"quadrant"`
	Ring     int8   `json:"ring" csv:"ring"`
}

// Read csv spec file and create Blips from that
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
	
	blips := Blips {
		Blips: smallBlips,
	}

	return blips
}

// Read json spec file and create Blips from that
func ReadJsonSpec(filePath string) Blips {
	// Open file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened json file!")
	defer jsonFile.Close()

	// Read file
	byteValue, _ := io.ReadAll(jsonFile)

	var blips Blips

	json.Unmarshal(byteValue, &blips)

	return blips
}
