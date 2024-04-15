package Merger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
)

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
	err = Verifier.DuplicateRemoval(filepaths...)
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
