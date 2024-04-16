package Fetcher

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
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
	defer os.Remove("specfile.txt")
	defer os.RemoveAll("./cache/")
	defer DotGitDelete()

	// dev repo link and create specfile
	url := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	//TODO: Change this to main once templates folder is on main
	branch := "main"
	data := []byte("examples/csv_templates/template.csv")
	os.WriteFile("./specfile.txt", data, 0644)
	specFile := "specfile.txt"

	err := FetchFiles(url, branch, specFile)

	if err != nil {
		t.Errorf("FetchFiles returned an err %v", err)
	}

	_, err = os.Stat("./cache/template.csv")
	if os.IsNotExist(err) {
		t.Errorf("File wasn't downloaded or wasn't moved correctly: %v", err)
	}

	file, err := os.Open("./cache/template.csv")
	if err != nil {
		t.Errorf("Failed to open template.csv. %v", err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	template_lines := []string{}
	for scanner.Scan() {
		template_lines = append(template_lines, scanner.Text())
	}

	expected_lines := []string{"name,ring,quadrant,isNew,move,description",
	"Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.",
	"web,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.",
	"react,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit."}

	for i := range expected_lines {
		if !(expected_lines[i] == template_lines[i]) {
			t.Errorf("Mismatch in downloaded file. Expected: %v \n Retrieved: %v",expected_lines[i], template_lines[i])
		}
	}
}
func TestListingReposForFetch(t *testing.T) {
	defer os.RemoveAll("./README.md")
	defer os.RemoveAll("./TestingTextFile.txt")
	defer os.RemoveAll("./cache/")
	//syscall.Rmdir(dirName)
	defer DotGitDelete()
	
	// Creates a text file named "TestingTextFile.txt" to hold a temporary
	// list of repo specifications for testing the ListingReposForFetch function
	textFile, errCreate := os.Create("./TestingTextFile.txt")
	if errCreate != nil {
		t.Errorf(errCreate.Error() + " : Couldnt create txt file for testing : TestListingReposForFetch")
	}

	_, errWrite := textFile.WriteString("/README.md")
	if errWrite != nil {
		t.Errorf(errWrite.Error() + " : Couldnt write to txt file for testing : TestListingReposForFetch")
	}

	url := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	branch := "main"
	specFile := "./TestingTextFile.txt"

	url2 := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	branch2 := "main"
	specFile2 := "./TestingTextFile.txt"

	url3 := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	branch3 := "main"
	specFile3 := "./TestingTextFile.txt"

	repos := []string {url, branch, specFile, url2, branch2, specFile2, url3, branch3, specFile3}

	err := ListingReposForFetch(repos)

	textFile.Close()

	//check if the error message is as expected.
	expectedErrorMessage := "failed at fetcher"
	if err != nil && strings.Contains(err.Error(), expectedErrorMessage) {
		os.RemoveAll("./TestingTextFile.txt")
		t.Errorf("Expected an error containing '%s', but got '%s'", expectedErrorMessage, err.Error())
	}

}

func TestGitDelete(t *testing.T) {
	// Makes a filepath for the folder
	dotGitpath := filepath.Join(".", ".git")

	// Makes a folder from the given parameters
	err := os.MkdirAll(dotGitpath, os.ModePerm)
	if err != nil {
		t.Errorf(err.Error() + " : Couldnt make the .git folder")
	}

	DotGitDelete()

	if _, err := os.Stat("./.git"); !os.IsNotExist(err) {
		t.Errorf(err.Error() + " : .git still exists")
	}
}
