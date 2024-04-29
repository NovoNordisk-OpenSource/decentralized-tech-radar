package Merger

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

var csvfile1 string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Infrastructure,false,2,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing`

var csvfile2 string = `name,ring,quadrant,isNew,moved,description
Python,Hold,Language,false,0,Its a programming Language
Visual Studio,Trial,Infrastructure,false,1,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing`

var correctMerge string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Infrastructure,false,2,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing
Python,Hold,Language,false,0,Its a programming Language`

var TestFiles []string = []string{"testFile1.csv", "testFile2.csv"}

func createCsvFiles() {
	err := os.WriteFile("testFile1.csv", []byte(csvfile1), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("testFile2.csv", []byte(csvfile2), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanUp() {
	os.Remove("testFile1.csv")
	os.Remove("testFile2.csv")
	os.Remove("Merged_file.csv")
	os.RemoveAll("./cache")
}

func TestGetHeader(t *testing.T) {
	createCsvFiles()
	defer cleanUp()
	correctHeader := "name,ring,quadrant,isNew,moved,description\n"
	header, err := getHeader("testFile1.csv")
	if err != nil {
		t.Fatalf("getHeader() gave an error: %v", err)
	}
	if string(header) != correctHeader {
		t.Errorf("Header does not match expected:\nGot: %s\nExpected: %s", string(header), correctHeader)
	}
}

func TestDuplicateDeletion(t *testing.T) {
	createCsvFiles()
	defer cleanUp()

	var buf bytes.Buffer
	ReadCsvData(&buf, "./testFile1.csv", "./testFile2.csv")

	csv1, err := os.ReadFile("./testFile1.csv")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(csv1), csvfile1) {
		t.Errorf("csvFile1 does not match expected output.\nExpected: %s \n Actual: %s", csvfile1, csv1)
	}

	csv2, err := os.ReadFile("./testFile2.csv")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(csv2), csvfile2) {
		t.Errorf("csvFile2 does not match expected output.\nExpected: %s \n Actual: %s", csvfile2, csv2)
	}

	correctString := `Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Infrastructure,false,2,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing
Python,Hold,Language,false,0,Its a programming Language
Visual Studio,Trial,Infrastructure,false,1,An IDE
`

	bufferString := buf.String()
	if bufferString != correctString {
		t.Errorf("Buffer doesn't contain the correct data.\nExpected: %s\n\nActual: %s", correctString, correctString)
	}
}

func TestMergeCSV(t *testing.T) {
	// Setup
	createCsvFiles()
	defer cleanUp()

	// Call function
	err := MergeCSV(TestFiles)
	if err != nil {
		t.Fatalf("MergeCSV gave an error: %v", err)
	}

	// Check that file exists
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Merged_file.csv was not found")
	}

	// Check that file is merged correctly
	content, err := os.ReadFile("Merged_file.csv")
	if err != nil {
		t.Fatalf("Merged_file.csv could not be read: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, correctMerge) {
		t.Errorf("Merged file doesn't contain the expected data\nContained:\n%s", contentStr)
	}
}

func TestMergeFromFolder(t *testing.T) {
	createCsvFiles()
	defer cleanUp()

	err := os.Mkdir("cache", 0700)
	if err != nil {
		t.Fatal(err)
	}
	os.Rename(TestFiles[0], "./cache/"+TestFiles[0])
	os.Rename(TestFiles[1], "./cache/"+TestFiles[1])

	err = MergeFromFolder("./cache")
	if err != nil {
		t.Fatal(err)
	}

	// Check that file exists
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Merged_file.csv was not found")
	}

	// Check that file is merged correctly
	content, err := os.ReadFile("Merged_file.csv")
	if err != nil {
		t.Fatalf("Merged_file.csv could not be read: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, correctMerge) {
		t.Errorf("Merged file doesn't contain the expected data\nContained:\n%s", contentStr)
	}
}
