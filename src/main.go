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
		if len(fetchArgs) < 3 || len(fetchArgs) % 3 != 0{
			fmt.Println(string(fetchArgs[0]) + " " + " Incorrect number of arguments for fetch command")
			os.Exit(1)
		}

		// Iterate through command-line arguments after flags
		// Go by every group of 3 arguments
		for i := 0; i < len(fetchArgs)-2; i += 3 {
			// Check for flag and enough arguments
			repo := Fetcher.Repo{fetchArgs[i], fetchArgs[i+1], fetchArgs[i+2]} // Extract repository details
			repos = append(repos, repo)
		}

		// Call the Fetcher package function to fetch files from all repositories
		err := Fetcher.ListingReposForFetch(repos)
		if err != nil {
			fmt.Println("Error fetching files:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *file != "" && *fetch == "" {
		specs := SpecReader.ReadCsvSpec(*file)
		html.GenerateHtml(specs)
	}
}
