package SpecReader

import (
	"os"
	"testing"
)

var csvfile = `name,ring,quadrant,isNew,moved,description
TestBlip1,Assess,languages & frameworks,true,1,This is a description
TestBlip2,Adopt,tool,false,0,Also a description
TestBlip3,Assess,languages & frameworks,true,1,This is a description
TestBlip4,Adopt,tool,false,0,Also a description`

func createFile() {
	err := os.WriteFile("testcsv.csv", []byte(csvfile), 0644)
	if err != nil {
		panic(err)
	}
}

func cleanUp() {
	os.Remove("testcsv.csv")
}

func TestCsvToString(t *testing.T) {
	createFile()
	defer cleanUp()

	csvString := CsvToString("testcsv.csv")
	if csvString != csvfile {
		t.Errorf("Read csv file is not the same as expected!\nExpected:\n%s\n\nActual:\n%s", csvfile, csvString)
	}
}