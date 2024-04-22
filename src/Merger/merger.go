package Merger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"slices"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
)

// Map of alternative names for the same blip
var alt_names = make(map[string]string) //{"golang":"Go","go-lang:Go","cpp":"C++","csharp":"C#","cs":"C#","python3":"Python","py":"Python"}

func getHeader(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []byte{}, err // Propagate error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	headerBytes := scanner.Bytes()
	headerBytes = append(headerBytes, []byte("\n")...)

	return headerBytes, nil
}

func readCsvContent(filepath string) ([]byte, error) {
	var fileBytes []byte

	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		return fileBytes, err // Propagate error
	}
	defer file.Close()

	// Read file line by line, skipping first line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		fileBytes = append(fileBytes, scanner.Bytes()...)
		fileBytes = append(fileBytes, []byte("\n")...) // Add newline between each line in the file, otherwise it's all on one line
	}

	return fileBytes, nil
}

func MergeFromFolder(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return errors.New("Folder does not exist or could not be found. \nError: " + err.Error())
	} else if err != nil {
		return err
	}

	cachedRepos, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	var cachePaths []string
	for _, repo := range cachedRepos {
		if filepath.Ext(repo.Name()) == ".csv" {
			cachePaths = append(cachePaths, filepath.Join(folderPath, repo.Name()))
		}
	}

	if len(cachePaths) == 0 {
		fmt.Println("There are currently no files in the cache.")
	}

	MergeCSV(cachePaths)

	return nil
}

func MergeCSV(filepaths []string) error {
	os.Remove("Merged_file.csv") // Remove file in case it already exists
	var buf bytes.Buffer

	// Run data verifier on files
	err := Verifier.Verifier(filepaths...)
	if err != nil {
		panic(err)
	}

	// Run duplicate removal on files
	err = DuplicateRemoval(filepaths...)
	if err != nil {
		panic(err)
	}

	// Add header to buffer
	header, err := getHeader(filepaths[0])
	if err != nil {
		return err // Propagate error
	}
	buf.Write(header)

	// Read file content and add to buffer
	for _, file := range filepaths {
		content, err := readCsvContent(file)
		if err != nil {
			return err
		}
		buf.Write(content)
	}

	// Write combined files to one file
	err = os.WriteFile("Merged_file.csv", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func DuplicateRemoval(filepaths ...string) error {
	// Map functions as a set (name -> ring)
	var set = make(map[string][]string)
	for _, filepath := range filepaths {

		// Create temp file to overwrite primary file
		tempfile, err := os.Create("tempfile.csv")
		if err != nil {
			panic(err)
		}

		defer os.RemoveAll(tempfile.Name())
		defer tempfile.Close()

		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)

		// Skip header
		scanner.Scan()
		tempfile.WriteString(scanner.Text() + "\n")

		for scanner.Scan() {
			line := scanner.Text()
			// Faster than splitting
			// Panic handler
			name := ""
			index := strings.IndexByte(line, ',')
			if index != -1 {
				name = line[:index]
			}

			duplicateRemoval(name, line, tempfile, set)

		}
		file.Close()
		tempfile.Close()
		err = os.Rename("tempfile.csv", filepath)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func duplicateRemoval(name, line string, tempfile *os.File, set map[string][]string) error {
	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	real_name := name
	if alt_names[name] != "" {
		//TODO: Figure out how to handle numbers in names
		name = alt_names[strings.ToLower(name)]
	}

	if set[name] != nil {
		// Skips the name + first comma and does the same forward search for next comma
		ring := line[len(real_name)+1 : strings.IndexByte(line[len(real_name)+1:], ',')+len(real_name)+1]
		if !(slices.Contains(set[name], ring)) {
			set[name] = append(set[name], ring)
			tempfile.WriteString(line + "\n")
		}
	} else {
		set[name] = append(set[name], line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		tempfile.WriteString(line + "\n")
	}
	// Overwrite filepath with tempfile (has the removed changes)
	return nil
}
