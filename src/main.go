package main

import (
	"flag"
	"fmt"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SpecReader"
)

func main() {
	file := flag.String("file", "", "This takes a path to a csv file/string")
	flag.Parse()

	var specs SpecReader.Blips
	
	// testing csv reader
	if *file != "" {
		specs = SpecReader.ReadCsvSpec(*file)
		for i := 0; i < len(specs.Blips); i++ {
			fmt.Printf("Tech name: %s\n\tQuadrant: %d\n\tRing: %d\n", 
				specs.Blips[i].Name, specs.Blips[i].Quadrant, specs.Blips[i].Ring)
		}
	}
}
