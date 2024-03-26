package Fetcher

import (
	"os"
	"testing"
)

// the test expects an error returned if the repo url, branch, and specFile is invalid
func TestFetchFilesInvalidArguments(t *testing.T) {
	// Invalid URL, branch, and specFile arguments.
	url := "https://invalid-url.com/nonexistent-repo"
	branch := "branch"
	specFile := "nonexistent-file.txt"

	err := FetchFiles(url, branch, specFile)

	// We expect an error since the arguments are invalid
	if err == nil {
		t.Error("FetchFiles did not return an error when given invalid arguments")
	}
}

func TestFetchFilesValidArguments(t *testing.T) {
    //TODO: Maybe this needs to be split into 2 tests
    
	// dev repo link and create specfile
	url := "https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev.git"
	//TODO: Change this to main once templates folder is on main
    branch := "feat_git_fetcher"
    data := []byte("examples/csv_templates/template.csv")
    os.WriteFile("./specfile.txt",data,0644)
    specFile := "specfile.txt"

    err := FetchFiles(url,branch,specFile)

    if err != nil {
        t.Errorf("FetchFiles returned an err %v", err)
    }

    _, err = os.Stat("./cache/template.csv")
   if os.IsNotExist(err) {
        t.Errorf("File wasn't downloaded or wasn't moved correctly: %v",err)
    } 

    //TODO: Maybe add a test for contents of CSV file?

    defer os.Remove("specfile.txt")
    defer os.RemoveAll("./cache")

}   
