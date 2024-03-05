package SpecReader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Code inspired from:
// https://tutorialedge.net/golang/parsing-json-with-golang/
type Blips struct {
	Blips []Blip `json:"Blips"`
}

type Blip struct {
	Name     string `json:"name"`
	Quadrant int8   `json:"quadrant"`
	Ring     int8   `json:"ring"`
}

// Read json spec file and create Blips from that
func ReadJsonSpec(filePath string) Blips {
	// Open file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File was opened!")

	// Read file
	byteValue, _ := io.ReadAll(jsonFile)

	var blips Blips

	json.Unmarshal(byteValue, &blips)

	// Clean up
	jsonFile.Close()

	return blips
}
