package Fetcher

import (
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