package main

import (
	"flag"
	"strings"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/Merger"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SpecReader"
	view "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/View"
)

func main() {
	file := flag.String("file", "", "This takes a path to a csv file/string")
	merge := flag.String("merge", "", "This takes paths to files that should be merged (space separated)")
	flag.Parse()

	var specs SpecReader.Blips
	// testing csv reader
	if *file != "" {
		specs = SpecReader.ReadCsvSpec(*file)
	}

	if *merge != "" {
		Merger.MergeCSV(strings.Split(*merge, " "))
	}
	
	view.GenerateHtml(specs)
}
