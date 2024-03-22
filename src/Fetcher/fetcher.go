package Fetcher

import (
	"fmt"
	"os"
	"os/exec"
)

type Repo struct {
	URL      string
	Branch   string
	SpecFile string
}

func FetchFiles(url, branch, specFile string) error {
	cmd := exec.Command("python", "./Fetcher/fetchfile.py", url, branch, specFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed at fetcher: %s. Error: %s", out, err)
	}

	errDotGit := DotGitDelete()
	if errDotGit != nil {
		return errDotGit
	}

	return nil
}

func ListingReposForFetch(repos []Repo) error {
	for _, repo := range repos {
		err := FetchFiles(repo.URL, repo.Branch, repo.SpecFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func DotGitDelete() error {
	// Check if the .git directory exists in the current directory
	if _, err := os.Stat("./.git"); err == nil {
		// Remove the .git directory if it exists (we are still in src which is the only place it'll remove the .git folder)
		err := os.RemoveAll("./.git")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}





// this is for testing the FetchFiles function

func FetchFilesTest(url, branch, specFile string) error {
	cmd := exec.Command("python", "./fetchfile.py", url, branch, specFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed at fetcher: %s. Error: %s", out, err)
	}

	errDotGit := DotGitDeleteTest()
	if errDotGit != nil {
		return errDotGit
	}

	return nil
}

func ListingReposForFetchTest(repos []Repo) error {
	for _, repo := range repos {
		err := FetchFilesTest(repo.URL, repo.Branch, repo.SpecFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func DotGitDeleteTest() error {
	// Check if the .git directory exists in the current directory
	if _, err := os.Stat("./Fetcher/.git"); err == nil {
		// Remove the .git directory if it exists (we are still in src which is the only place it'll remove the .git folder)
		err := os.RemoveAll("./Fetcher/.git")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}