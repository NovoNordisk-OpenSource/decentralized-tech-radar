package main


import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Fetcher"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Merger"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
	html "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/HTML"
)

func main() {
	var repos []Fetcher.Repo
  
	file := flag.String("file", "", "This takes a path to a csv file/string")
	fetch := flag.String("fetch", "", "This command will fetch a file from a git repo")
	merge := flag.String("merge", "", "This takes paths to files that should be merged (space separated)")
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
	

	if *merge != "" {
		err := Merger.MergeCSV(strings.Split(*merge, " "))
		if err != nil {
			panic(err)
		}
	}

	if *file != "" && *fetch == "" {
		specs := SpecReader.ReadCsvSpec(*file)
		if true == false {
			print(specs)
		}
		csvString := SpecReader.CsvToString(*file)
		html.GenerateHtml(/*specs, */csvString)
	}
}
