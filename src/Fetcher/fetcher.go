package Fetcher

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
)

func FetchFiles(url, branch, specFile string) error {
	defer DotGitDelete()

	// Create cache dir if it doesn't exist
	if _, err := os.Stat("cache"); os.IsNotExist(err) {
		err := os.Mkdir("cache", 0700)
		errHandler(err)
	}

	// Pulls files and returns the paths to said files
	seenFolders := make(map[string]string)
	paths, err := puller(url, branch, specFile)
	if err != nil {
		return err
	}

	for _, path := range paths {
	
    var fileName []string 
    if runtime.GOOS == "windows" {
      fileName = strings.Split(path, "\\")
    } else {
      fileName = strings.Split(path, "/")
    }
		
		if _, ok := seenFolders[fileName[0]]; !ok {
			seenFolders[fileName[0]] = ""
		}

		
		os.Rename(path, ("cache/" + fileName[len(fileName)-1]))
		
		// Runs verifier on downloaded file to remove duplicates and ensure data integrity
		
		file := "./cache/"+fileName[len(fileName)-1]
		err := Verifier.Verifier(file)
		if err != nil {
			fmt.Printf("File has problems not correctly formatted CSV file: "+file +"\ncontinuing to next file")
		}
	}

	for folder, _ := range seenFolders {
        if folder != "cache" {
            os.RemoveAll(("./" + folder))
        }
    }

	return nil
}

func ListingReposForFetch(repos []string) error {
	for i := 0; i < len(repos); i += 3 {
		err := FetchFiles(repos[i], repos[i+1], repos[i+2])
		if err != nil {
			return err
		}
	}
	return nil
}

func DotGitDelete() {
	// Check if the .git directory exists in the current directory
	if _, err := os.Stat("./.git"); err == nil {
		// Remove the .git directory if it exists (we are still in src which is the only place it'll remove the .git folder)
		os.RemoveAll("./.git")
	}
}

func errHandler(err error, params ...string) {
	if err != nil {
		panic(err.Error() + strings.Join(params, " "))
	}
}

func executer(cmd *exec.Cmd) error {
	//TODO: Figure out a way to take strings as input and build cmd
	_, err := cmd.CombinedOutput()

	return err
}

func puller(url, branch, specFile string) ([]string, error) {
	paths := []string{}

	// Create dummy repo
	cmd := exec.Command("git", "init")
	err := executer(cmd)
	if err != nil {
		return paths, err
	}

	//Enable sparse Checkout
	cmd = exec.Command("git", "config", "core.sparseCheckout", "true")
	err = executer(cmd)
	if err != nil {
		return paths, err
	}

	// Add whitelist to sparse-checkout
	fileData, err := os.ReadFile(specFile)
	err = os.WriteFile(".git/info/sparse-checkout", fileData, 0644)
	if err != nil {
		return paths, err
	}

	// Add remote repo
	cmd = exec.Command("git", "remote", "add", "origin", url)
	err = executer(cmd)
	if err != nil {
		return paths, err
	}

	// git pull from remote repo
	cmd = exec.Command("git", "pull", "origin", branch, "--depth=1")
	err = executer(cmd)
	if err != nil {
		return paths, err
	}

	//remove .git folder
	err = os.RemoveAll("./.git")
	if err != nil {
		return paths, err
	}

	// https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
	// This function recursively walks the directors inside the workdir and checks for csv files
	// These then get added to the cache later
	filepath.WalkDir(".", func(str string, dir fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		path_seg := strings.Split(str, "/")
		if path_seg[0] != "cache" {
			if filepath.Ext(dir.Name()) == ".csv" {
				paths = append(paths, str)
			}
		}
		return nil
	})

	return paths, nil
}