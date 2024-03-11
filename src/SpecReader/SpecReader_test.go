package SpecReader

import (
	"os"
	"testing"

	"github.com/gocarina/gocsv"
)

// Set up common test data
var testFileName string = "ForTesting"
var blip1 Blip = Blip{
	Name:     "IAmInQuadrant3Ring2",
	Quadrant: "Language",
	Ring:     "Assess",
	IsNew: true,
	Moved: 0,
	Description: "This is description",
}
var blip2 Blip = Blip{
	Name:     "IAmInQuadrant2Ring0",
	Quadrant: "Tool",
	Ring:     "Hold",
	IsNew: true,
	Moved: 1,
	Description: "This is description also",
}
var blip3 Blip = Blip{
	Name:     "IAmInQuadrant1Ring3",
	Quadrant: "Language",
	Ring:     "Adopt",
	IsNew: false,
	Moved: 2,
	Description: "This is description not",
}
var testBlips Blips = Blips{
	Blips: []Blip{blip1, blip2, blip3},
}

func createCsvFile() {
	csvFile, err := os.Create(testFileName + ".csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	err = gocsv.MarshalFile(testBlips.Blips, csvFile)
	if err != nil {
		panic(err)
	}
}

func CleanUp() {
	os.Remove(testFileName + ".csv")
}

func TestReadCsvSpec(t *testing.T) {
	// Set up
	createCsvFile()
	defer CleanUp()

	// Run test
	gottenBlips := ReadCsvSpec(testFileName + ".csv")
	
	// Check correctness
	for i := 0; i < len(testBlips.Blips); i++ {
		if gottenBlips.Blips[i] != testBlips.Blips[i] {
			t.Error("Read blips differs from expected")
		}
	}
}