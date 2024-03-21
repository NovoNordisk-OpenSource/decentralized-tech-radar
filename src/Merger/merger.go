package Merger

import (
	"bytes"
	"os"
)

func MergeCSV(file1 string, file2 string) {
	var buf bytes.Buffer
	b1, err := os.ReadFile(file1)
	if err != nil {
		panic(err)
	}
	buf.Write(b1)
	b2, err := os.ReadFile(file2)
	if err != nil {
		panic(err)
	}
	buf.Write(b2)
	err = os.WriteFile("Merged_file.csv", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}