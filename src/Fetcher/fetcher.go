package Fetcher

import (
	"os"
	"os/exec"
)

func FetchFiles(url, branch, specFile string) { 

    // Check if the .git directory exists in the current directory
    if _, err := os.Stat("./.git"); err == nil {
        // Remove the .git directory if it exists (we are still in src which is the only place it'll remove the .git folder)
        err := os.RemoveAll("./.git")
        if err != nil {
            panic("Failed to remove .git directory: " + err.Error())
        }
    }

    cmd := exec.Command("python", "./Fetcher/fetchfile.py", url, branch, specFile)
    out, err := cmd.CombinedOutput()
    if err != nil {
        panic(string(out) + " Failed at fetcher")
    }
 
}

