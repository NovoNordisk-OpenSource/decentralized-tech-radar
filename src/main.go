package main

import (
	"flag"
	"fmt"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SpecReader"
)

func main() {
	json := flag.String("json", "", "This takes a path to a json file/string")
	csv := flag.String("csv", "", "This takes a path to a csv file/string")
	flag.Parse()

	var specs SpecReader.Blips
	// testing json reader
	if *json != "" {
		specs = SpecReader.ReadJsonSpec(*json)
	
		for i := 0; i < len(specs.Blips); i++ {
			fmt.Printf("Tech name: %s\n\tQuadrant: %d\n\tRing: %d\n", 
				specs.Blips[i].Name, specs.Blips[i].Quadrant, specs.Blips[i].Ring)
		}
	}
	
	// testing csv reader
	if *csv != "" {
		specs = SpecReader.ReadCsvSpec(*csv)
		for i := 0; i < len(specs.Blips); i++ {
			fmt.Printf("Tech name: %s\n\tQuadrant: %d\n\tRing: %d\n", 
				specs.Blips[i].Name, specs.Blips[i].Quadrant, specs.Blips[i].Ring)
		}
	}
}
