package main

import  (
	"flag"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/specReader"
	view "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/View"
)

func main() {
	file := flag.String("file", "", "This takes a path to a csv file/string")
	flag.Parse()

	var specs SpecReader.Blips
	// testing csv reader
	if *file != "" {
		specs = SpecReader.ReadCsvSpec(*file)
	} else {
		panic("No file was given (oh no)")
	}
	
	view.GenerateHtml(specs)
}
