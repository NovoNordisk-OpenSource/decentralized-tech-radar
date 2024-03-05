package SpecReader

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
)

// Set up common test data
var testFileName string = "ForTesting"
var blip1 Blip = Blip{
	Name:     "IAmInQuadrant3Ring2",
	Quadrant: 3,
	Ring:     2,
}
var blip2 Blip = Blip{
	Name:     "IAmInQuadrant2Ring0",
	Quadrant: 2,
	Ring:     0,
}
var blip3 Blip = Blip{
	Name:     "IAmInQuadrant1Ring3",
	Quadrant: 1,
	Ring:     3,
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

func createJsonFile() {
	jsonString, err := json.Marshal(testBlips)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(testFileName + ".json", jsonString, 0666)
	if err != nil {
		panic(err)
	}
}

func CleanUp() {
	os.Remove(testFileName + ".json")
	os.Remove(testFileName + ".csv")
}

func TestReadJsonSpec(t *testing.T) {
	// Set up
	createJsonFile()
	defer CleanUp()

	// Run function
	gottenBlips := ReadJsonSpec(testFileName + ".json")

	// Check correctness 
	for i := 0; i < len(testBlips.Blips); i++ {
		currentTestBlip := testBlips.Blips[i]
		currentGottenBlip := gottenBlips.Blips[i]
		nameMatch := currentGottenBlip.Name == currentTestBlip.Name
		quadrantMatch := currentGottenBlip.Quadrant == currentTestBlip.Quadrant
		ringMatch := currentGottenBlip.Ring == currentTestBlip.Ring
		if !(nameMatch && quadrantMatch && ringMatch) {
			t.Error("Read blips differs from expected")
		} 
	}
}

func TestReadCsvSpec(t *testing.T) {
	// Set up
	createCsvFile()
	defer CleanUp()

	// Run test
	gottenBlips := ReadCsvSpec(testFileName + ".csv")
	
	// Check correctness
	for i := 0; i < len(testBlips.Blips); i++ {
		currentTestBlip := testBlips.Blips[i]
		currentGottenBlip := gottenBlips.Blips[i]
		nameMatch := currentGottenBlip.Name == currentTestBlip.Name
		quadrantMatch := currentGottenBlip.Quadrant == currentTestBlip.Quadrant
		ringMatch := currentGottenBlip.Ring == currentTestBlip.Ring
		if !(nameMatch && quadrantMatch && ringMatch) {
			t.Error("Read blips differs from expected")
		} 
	}
}