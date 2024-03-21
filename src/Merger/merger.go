package Merger

import (
	"bytes"
	"os"
)

func MergeCSV(file1 string, file2 string) {
	os.Remove("Merged_file.csv")
	var buf bytes.Buffer
	b1, err := os.ReadFile(file1)
	if err != nil {
		panic(err)
	}
	buf.Write(b1)
	buf.Write([]byte("\n"))
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