package SpecReader

import (
	"os"
)

func CsvToString(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// Build the CSV string
	csvString := string(data)

	return csvString
}
