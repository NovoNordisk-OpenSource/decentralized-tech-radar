package Merger

import (
	"log"
	"os"
	"strings"
	"testing"
)

var csvfile1 string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Tool,false,2,An IDE
Dagger IO,Assess,Tool,true,1,Its a workflow thing`

var csvfile2 string = `name,ring,quadrant,isNew,moved,description
Python,Halt,Language,false,0,Its a programming Language
Visual Studio,Trial,Tool,false,1,An IDE
Dagger IO,Assess,Tool,true,1,Its a workflow thing`

var correctMerge string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Tool,false,2,An IDE
Dagger IO,Assess,Tool,true,1,Its a workflow thing
Python,Halt,Language,false,0,Its a programming Language
Visual Studio,Trial,Tool,false,1,An IDE
Dagger IO,Assess,Tool,true,1,Its a workflow thing`

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

func TestReadCsvContent(t *testing.T) {
	createCsvFiles()
	defer cleanUp()

	correctContent := `Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Tool,false,2,An IDE
Dagger IO,Assess,Tool,true,1,Its a workflow thing
`
	readContent, err := readCsvContent("testFile1.csv")
	if err != nil {
		t.Fatalf("readCsvContent() gave an error: %v", err)
	}
	if string(readContent) != correctContent {
		t.Errorf("Read content does not match expected:\nGot:\n\t%s\nExpected:\n\t%s", string(readContent), correctContent)
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