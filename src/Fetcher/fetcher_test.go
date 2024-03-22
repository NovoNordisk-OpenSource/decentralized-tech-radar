package Fetcher

import (
	"strings"
	"testing"
    "os"
    "path/filepath"
)

// the test expects an error returned if the repo url, branch, and specFile is invalid
func TestFetchFilesInvalidArguments(t *testing.T) {
	// Invalid URL, branch, and specFile arguments.
	url := "https://invalid-url.com/nonexistent-repo"
	branch := "branch"
	specFile := "nonexistent-file.txt"

	err := FetchFilesTest(url, branch, specFile)

	// We expect an error since the arguments are invalid
	if err == nil {
		t.Error("FetchFiles did not return an error when given invalid arguments")
	} else {
		//check if the error message is as expected.
		expectedErrorMessage := "failed at fetcher"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Expected an error containing '%s', but got '%s'", expectedErrorMessage, err.Error())
		}
	}
}

func TestListingReposForFetch(t *testing.T) {
    // Creates a txt file for testing the  
    textFile, errCreate := os.Create("./TestingTextFile.txt")
    if errCreate != nil {
        t.Errorf(errCreate.Error() + " : Couldnt create txt file for testing : TestListingReposForFetch")
    } 
    
    _, errWrite := textFile.WriteString("README.md")
    if errWrite != nil {
        t.Errorf(errWrite.Error() + " : Couldnt write to txt file for testing : TestListingReposForFetch")
    }

    // github
    url := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	branch := "main"
	specFile := "./TestingTextFile.txt"

    // gitlab
    url2 := "https://gitlab.com/nagyv-gitlab/gitops-test"
	branch2 := "master"
	specFile2 := "./TestingTextFile.txt"

    // bitbucket
    url3 := "https://bitbucket.org/tech-radar/tech-radar/src/main/"
	branch3 := "main"
	specFile3 := "./TestingTextFile.txt"

    var repos []Repo
    repo := Repo{url, branch, specFile}
    repo2 := Repo{url2, branch2, specFile2}
    repo3 := Repo{url3, branch3, specFile3}
    
    repos = append(repos, repo, repo2, repo3)

    err := ListingReposForFetchTest(repos)
    textFile.Close()

    //check if the error message is as expected.
    expectedErrorMessage := "failed at fetcher"
    if  err != nil && strings.Contains(err.Error(), expectedErrorMessage) {
        os.RemoveAll("./TestingTextFile.txt")
        t.Errorf("Expected an error containing '%s', but got '%s'", expectedErrorMessage, err.Error())
    }
    os.RemoveAll("./TestingTextFile.txt")
    os.RemoveAll("./README.md")
}

func TestGitDelete(t *testing.T) {
    // Makes a filepath for the folder
    dotGitpath := filepath.Join(".", ".git")
    
    // Makes a folder from the given parameters
    err := os.MkdirAll(dotGitpath, os.ModePerm)
    if err != nil {
        t.Errorf(err.Error() + " : Couldnt make the .git folder")
    }

    errDot := DotGitDelete()
    if errDot != nil {
        t.Errorf("DotGitDelete could not delete the folder")
    }

    if _, err := os.Stat("./.git"); !os.IsNotExist(err) {
        t.Errorf(err.Error() + " : .git still exists")
    }
}
