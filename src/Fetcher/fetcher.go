package Fetcher

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
)

var token sync.Mutex
var finished int

func FetchFiles(url, branch, specFile string, ch chan error ) {
	// Pulls files and returns the paths to said files
	seenFolders := make(map[string]string)
	paths, err := puller(url, branch, specFile)
	if err != nil {
		ch <- err
	}
	for _, path := range paths {
		var fileNamePath []string
		if runtime.GOOS == "windows" {
			fileNamePath = strings.Split(path, "\\")
		} else {
			fileNamePath = strings.Split(path, "/")
		}

		if _, ok := seenFolders[fileNamePath[0]]; !ok {
			seenFolders[fileNamePath[0]] = ""
		}

		// Handle renaming even when two files have the same name, by adding a number to the end
		i := 0
		var newFileName string
		token.Lock()
		for {
			fileName := fileNamePath[len(fileNamePath)-1]
			if i == 0 { // Don't add a 0 to filename
				newFileName = "cache/" + fileName
			} else {
				newFileName = fmt.Sprintf("cache/" + fileName[:len(fileName)-4] + "(%d)" + fileName[len(fileName)-4:], i)
			}

			if _, err := os.Stat(newFileName); os.IsNotExist(err) {
				err := os.Rename(path, newFileName)
				if err != nil {
					panic(err)
				}
				token.Unlock()
				break
			}
			i++
		}
		
		// Runs data integrity verifier on downloaded file
		// file := "./cache/"+fileName[len(fileName)-1]
		err = Verifier.Verifier(newFileName)
		if err != nil {
			fmt.Printf("CSV file contains incorrectly formatted content: "+newFileName +"\nContinuing to next file...")
		}
	}

	ch <- nil
}

func ListingReposForFetch(repos []string) error {
	// Create cache dir if it doesn't exist
	if _, err := os.Stat("cache"); os.IsNotExist(err) {
		err := os.Mkdir("cache", 0700)
		errHandler(err)
	}
	
	// Create temp folder for .git folders
	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		err := os.Mkdir("temp", 0700)
		errHandler(err)
	}
	defer os.RemoveAll("temp")
	
	channel := make(chan error)
	for i := 0; i < len(repos); i += 3 {
		go FetchFiles(repos[i], repos[i+1], repos[i+2], channel)
	}
	finished = 0
	go progressBar(len(repos)/3)
	for i := 0; i < len(repos)/3; i++ {
		err := <- channel
		if err != nil {
			return err
		}
		finished++
	}
	return nil
}

func progressBar(numOfFiles int) {
	progressBar := []string{}
	for i := 0; i < numOfFiles; i++ {
		progressBar = append(progressBar, ".")
	}
	for {
		for _, r := range `-\|/` {
			for i := 0; i < finished; i++ {
				progressBar[i] = "#"
			}
			percent_fin := float32(finished) / float32(numOfFiles) * 100.0
			fmt.Printf("\r%c [%s] %d%%", r, strings.Join(progressBar, ""), int32(percent_fin))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func errHandler(err error, params ...string) {
	if err != nil {
		panic(err.Error() + strings.Join(params, " "))
	}
}

func executer(cmd *exec.Cmd, folder string) error {
	//TODO: Figure out a way to take strings as input and build cmd
	cmd.Dir = folder
	_, err := cmd.CombinedOutput()

	return err
}

func puller(url, branch, specFile string) ([]string, error) {
	paths := []string{}
	
	// Create temp folder for git in the system temp folder
	var randomNum int
	var tempFolder string
	for {
		randomNum = rand.Int()
		tempFolder = fmt.Sprintf("temp/%d", randomNum)
		if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
			os.Mkdir(tempFolder, 0700)
			break
		}
	}

	// Create dummy repo
	cmd := exec.Command("git", "init")
	err := executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	//Enable sparse Checkout
	cmd = exec.Command("git", "config", "core.sparseCheckout", "true")
	err = executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	// Add whitelist to sparse-checkout
	fileData, err := os.ReadFile(specFile)
	if err != nil {
		return paths, err
	}

	err = os.WriteFile(tempFolder + "/.git/info/sparse-checkout", fileData, 0644)
	if err != nil {
		return paths, err
	}

	// Add remote repo
	cmd = exec.Command("git", "remote", "add", "origin", url)
	err = executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	// git pull from remote repo
	cmd = exec.Command("git", "pull", "origin", branch, "--depth=1")
	err = executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	// https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
	// This function recursively walks the directors inside the workdir and checks for csv files
	// These then get added to the cache later
	filepath.WalkDir(tempFolder, func(str string, dir fs.DirEntry, e error) error {
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