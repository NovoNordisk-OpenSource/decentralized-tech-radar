package main

import (
	"flag"
	"fmt"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SpecReader"
)

func main() {
	name := flag.String("name", "world", "This takes a name/string ")
	flag.Parse()

	specs := SpecReader.ReadJsonSpec(*name)

	for i := 0; i < len(specs.Blips); i++ {
		fmt.Printf("Tech name: %s\n\tQuadrant: %d\n\tRing: %d\n", 
			specs.Blips[i].Name, specs.Blips[i].Quadrant, specs.Blips[i].Ring)
	}
}
