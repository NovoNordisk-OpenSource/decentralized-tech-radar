package SpecReader

import (
	"encoding/json"
	"os"
	"testing"
)

var testFileName string = "ForTesting"
var testBlips Blips

func createJsonFile() {
	blip1 := Blip{
		Name:     "IAmInQuadrant3Ring2",
		Quadrant: 3,
		Ring:     2,
	}
	blip2 := Blip{
		Name:     "IAmInQuadrant2Ring0",
		Quadrant: 2,
		Ring:     0,
	}
	blip3 := Blip{
		Name:     "IAmInQuadrant1Ring3",
		Quadrant: 1,
		Ring:     3,
	}
	testBlips = Blips{
		Blips: []Blip{blip1, blip2, blip3},
	}

	jsonString, err := json.Marshal(testBlips)
	if err != nil {
		panic(err)
	}

	os.WriteFile(testFileName + ".json", jsonString, 0666)
}

func CleanUp() {
	err := os.Remove(testFileName + ".json")
	if err != nil {
		panic(err)
	}
}

func TestReadJsonSpec(t *testing.T) {
	// Setup
	createJsonFile()
	defer CleanUp()

	// Run test
	gottenBlips := ReadJsonSpec(testFileName + ".json")
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
