package main

import  (
	"flag"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
	html "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/HTML"
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
	
	html.GenerateHtml(specs)
}
