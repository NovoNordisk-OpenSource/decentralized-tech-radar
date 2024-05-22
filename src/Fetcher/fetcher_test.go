package Fetcher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// the test expects an error returned if the repo url, branch, and specFile is invalid
func TestFetchFilesInvalidArguments(t *testing.T) {
	// Invalid URL, branch, and specFile arguments.
	url := "https://invalid-url.com/nonexistent-repo"
	branch := "branch"
	specFile := "nonexistent-file.txt"
	
	os.Mkdir("cache", 0700)
	defer os.RemoveAll("cache")
	os.Mkdir("temp", 0700)
	defer os.RemoveAll("temp")
	ch := make(chan error)
	go FetchFiles(url, branch, specFile, ch)
	err := <- ch
	// We expect an error since the arguments are invalid
	if err == nil {
		t.Error("FetchFiles did not return an error when given invalid arguments")
	}
}

func TestFetchFilesValidArguments(t *testing.T) {
	defer os.Remove("specfile.txt")
	defer os.RemoveAll("./cache/")

	// dev repo link and create specfile
	url := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	//TODO: Change this to main once templates folder is on main
	branch := "main"
	data := []byte("examples/csv_templates/template.csv")
	os.WriteFile("./specfile.txt", data, 0644)
	specFile := "specfile.txt"
	
	os.Mkdir("cache", 0700)
	os.Mkdir("temp", 0700)
	defer os.RemoveAll("temp")
	ch := make(chan error)
	go FetchFiles(url, branch, specFile, ch)
	err := <- ch
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

	
	//TODO: This really should have been done using a variable
	expected_lines := []string{"name,ring,quadrant,isNew,moved,description",
	fmt.Sprintf("Python,hold,Languages & Frameworks,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.<br>Repos:<br> <a href=%s>decentralized-tech-radar</a>", url),
	fmt.Sprintf("web,hold,Languages & Frameworks,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.<br>Repos:<br> <a href=%s>decentralized-tech-radar</a>", url),
	fmt.Sprintf("react,hold,Languages & Frameworks,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.<br>Repos:<br> <a href=%s>decentralized-tech-radar</a>", url)}

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
