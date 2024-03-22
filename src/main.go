package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/Fetcher"
	html "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/HTML"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
)

func main() {
	var fetchFlag string
	var repos []Fetcher.Repo
	
	fetchHelp := "This command will fetch files from a list of git repos (These are separated by semicolons, each repo with URL branch specFile).\n \nHere is an example\n \ngo run main.go -fetch https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev main ./Fetcher/something.txt -fetch https://github.com/JonasSkjodt/CopenhagenBuzz master ./Fetcher/something.txt"
	flag.StringVar(&fetchFlag, "fetch", "", fetchHelp)
	file := flag.String("file", "", "This takes a path to a csv file/string")
	flag.Parse()

	if fetchFlag == "" {
		fmt.Println("No fetch command was given")
		os.Exit(1)
	}

	// Iterate through command-line arguments after flags
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-fetch" && i+3 < len(os.Args) { // Check for flag and enough arguments
		  repo := Fetcher.Repo{os.Args[i+1], os.Args[i+2], os.Args[i+3]} // Extract repository details
		  repos = append(repos, repo)
		  i += 3 // Skip to the next potential flag or argument
		}
	  }
	
	  if len(repos) == 0 {
		fmt.Println("No fetch commands were given")
		os.Exit(1)
	  }
	
	  // Call the Fetcher package function to fetch files from all repositories
	  err := Fetcher.FetchAllRepos(repos)
	  if err != nil {
		fmt.Println("Error fetching files:", err)
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
