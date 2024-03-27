package Merger

import (
	"bufio"
	"bytes"
	"os"
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

func MergeCSV(filepaths []string) error {
	os.Remove("Merged_file.csv") // Remove file in case it already exists
	var buf bytes.Buffer

	// Add header to buffer
	header, err := getHeader(filepaths[0])
	if err != nil{
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