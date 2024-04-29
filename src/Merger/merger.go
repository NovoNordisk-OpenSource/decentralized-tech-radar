package Merger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

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

	// Run data verifier on files
	err := Verifier.Verifier(filepaths...)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	// Add header to buffer
	header, err := getHeader(filepaths[0])
	if err != nil {
		return err // Propagate error
	}
	buf.Write(header)

	// Read csv data which removes duplicates
	// This only adds non-duplicates to the buffer
	err = ReadCsvData(&buf, filepaths...)
	if err != nil {
		panic(err)
	}

	// Write combined files to one file
	err = os.WriteFile("Merged_file.csv", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadCsvData(buffer *bytes.Buffer, filepaths ...string) error {
	// Map functions as a set (name -> quadrant)
	var set = make(map[string][]string)
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		scanFile(file, buffer, set)
	}
	return nil
}

func scanFile(file *os.File, buffer *bytes.Buffer, set map[string][]string) {
	scanner := bufio.NewScanner(file)

	// Skip header
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		// Faster than splitting
		// Panic handler
		name := ""
		index := strings.IndexByte(line, ',')
		if index != -1 {
			name = line[:index]
		}

		duplicateRemoval(name, line, buffer, set)
	}
}

func duplicateRemoval(name, line string, buffer *bytes.Buffer, set map[string][]string) error {
	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	real_name := name
	if alt_names[name] != "" {
		//TODO: Figure out how to handle numbers in names
		name = alt_names[strings.ToLower(name)]
	}

	ring_len := len(line[len(real_name)+1 : strings.IndexByte(line[len(real_name)+1:], ',')+len(real_name)+1])
	if set[name] != nil {
	// Skips the name + ring + 2 commas and does the same forward search for next comma
	// Example of a line from a csv file:
	// 		Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.
	// Quadrant:
	//		language
		quadrant := line[len(real_name)+ring_len+2 : strings.IndexByte(line[len(real_name)+ring_len+2:], ',')+len(real_name)+ring_len+2]
		if !(slices.Contains(set[name], quadrant)) {
			set[name] = append(set[name], quadrant)
			buffer.Write([]byte(line + "\n"))
		}
	} else {
		set[name] = append(set[name], line[len(name)+ring_len+2:strings.IndexByte(line[len(name)+ring_len+2:], ',')+len(name)+ring_len+2])
		buffer.Write([]byte(line + "\n"))
	}

	return nil
}
