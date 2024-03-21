package main

import  (
	"strings"
	"flag"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/Fetcher"
	html "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/HTML"
)

func main() {
	file := flag.String("file", "", "This takes a path to a csv file/string")
	fetch := flag.String("fetch", "", "This command will fetch a file from a git repo")
	flag.Parse()

	if *fetch != "" {
		fetchArgs := strings.Split(*fetch, " ")
		Fetcher.FetchFiles(fetchArgs[0], fetchArgs[1])
	} else {
		panic("No file was given (oh no)")
	}
	
	var specs SpecReader.Blips
	// testing csv reader
	if *file != "" {
		specs = SpecReader.ReadCsvSpec(*file)
	} else {
		panic("No file was given (oh no)")
	}
	
	html.GenerateHtml(specs)
}
