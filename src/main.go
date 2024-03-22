package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/Fetcher"
	html "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/HTML"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
)

func main() {
	file := flag.String("file", "", "This takes a path to a csv file/string")
	fetch := flag.String("fetch", "", "This command will fetch a file from a git repo")
	flag.Parse()

	/*if *fetch != "" {
		fetchArgs := strings.Split(*fetch, " ")
		Fetcher.FetchFiles(fetchArgs[0], fetchArgs[1], fetchArgs[2])
	} else {
		panic("No file was given (oh no)")
	}*/
	if *fetch != "" {
        fetchArgs := strings.Split(*fetch, " ")
        if len(fetchArgs) != 3 {
            fmt.Println("Incorrect number of arguments for fetch command")
            os.Exit(1)
        }
        err := Fetcher.FetchFiles(fetchArgs[0], fetchArgs[1], fetchArgs[2])
        if err != nil {
            fmt.Println("Error fetching files:", err)
            os.Exit(1)
        }
    } else {
        fmt.Println("No fetch command was given")
        os.Exit(1)
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
