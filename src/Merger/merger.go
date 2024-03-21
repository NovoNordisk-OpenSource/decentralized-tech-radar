package Merger

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func readCsvContent(filepath string) []byte {
	var fileBytes []byte

	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// Read file line by line, skipping first line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		fileBytes = append(fileBytes, scanner.Bytes()...)
		fileBytes = append(fileBytes, []byte("\n")...)
	}

	return fileBytes
}

func MergeCSV(filepaths []string, header string) {
	os.Remove("Merged_file.csv")
	var buf bytes.Buffer

	// Add header to buffer
	buf.Write([]byte(header + "\n"))

	// Read file content and add to buffer
	for _, file := range filepaths {
		buf.Write(readCsvContent(file))
	}

	// Write combined files to one file
	err := os.WriteFile("Merged_file.csv", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}