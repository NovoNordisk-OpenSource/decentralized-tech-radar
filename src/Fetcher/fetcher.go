package Fetcher

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func FetchFiles(url, branch, specFile string) error {

	// Check if the .git directory exists in the current directory
	if _, err := os.Stat("./.git"); err == nil {
		// Remove the .git directory if it exists (we are still in src which is the only place it'll remove the .git folder)
		err := os.RemoveAll("./.git")
		if err != nil {
			return err
		}
	}

	puller(url,branch,specFile)

	return nil
}

func errHandler(err error, params ...string) {
	if err != nil {
		panic(err.Error() + strings.Join(params, " "))
	}
}

func executer(cmd *exec.Cmd) {
	//TODO: Figure out a way to take strings as input and build cmd
	out, err := cmd.CombinedOutput()
	errHandler(err, string(out))
}

// TODO: should probably return the path of the file it downloaded
func puller(url, branch, specFile string) []string {
	// Create dummy repo
	cmd := exec.Command("git", "init")
	executer(cmd)

	//Enable sparse Checkout
	cmd = exec.Command("git", "config", "core.sparseCheckout", "true")
	executer(cmd)

	// Add whitelist to sparse-checkout
	fileData, err := os.ReadFile(specFile)
	errHandler(err)
	err = os.WriteFile(".git/info/sparse-checkout", fileData, 0644)
	errHandler(err)

	// Add remote repo
	cmd = exec.Command("git", "remote", "add", "origin", url)
	executer(cmd)

	// git pull from remote repo
	cmd = exec.Command("git", "pull", "origin", branch)
	executer(cmd)

	//remove .git folder
	err = os.RemoveAll("./.git")
	errHandler(err)

	// https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
	paths := []string{}
	filepath.WalkDir(".", func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ".csv" {
			paths = append(paths, s)
		}
		return nil
	})

	return paths

}
