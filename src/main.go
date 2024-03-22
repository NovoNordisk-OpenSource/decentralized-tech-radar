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
	var repos []Fetcher.Repo
	
	fetch := flag.String("fetch", "", "This command will fetch files from a list of git repos.\n \nHere is an example\n \n go run main.go --fetch \"https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev main ./Fetcher/something.txt https://github.com/JonasSkjodt/CopenhagenBuzz master ./Fetcher/something.txt https://gitlab.com/nagyv-gitlab/gitops-test master ./Fetcher/something.txt\"")
	file := flag.String("file", "", "This takes a path to a csv file/string")
	flag.Parse()

	if *fetch != "" {
		fetchArgs := strings.Split(*fetch, " ")
		if len(fetchArgs) < 3 {
			fmt.Println(string(fetchArgs[0]) + " " + " Incorrect number of arguments for fetch command")
			os.Exit(1)
		}

		// Iterate through command-line arguments after flags
		for i := 0; i < len(fetchArgs); i++ {
			// Check for flag and enough arguments
			repo := Fetcher.Repo{fetchArgs[i], fetchArgs[i+1], fetchArgs[i+2]} // Extract repository details
			repos = append(repos, repo)
			i += 2 // Skip to the next potential flag or argument
		}

		// Call the Fetcher package function to fetch files from all repositories
		err := Fetcher.ListingReposForFetch(repos)
		if err != nil {
			fmt.Println("Error fetching files:", err)
			os.Exit(1)
		}
	}

	if *file != "" && *fetch == "" {
		var specs SpecReader.Blips
		// testing csv reader
		if *file != "" {
			specs = SpecReader.ReadCsvSpec(*file)
		} else {
			panic("No file was given (oh no)")
		}

		html.GenerateHtml(specs)
	}
}