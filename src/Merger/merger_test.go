package Merger

import (
	"bytes"
	"fmt"
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

// The strings below are used exclusively to test
// ReadCsvFile's majority-vote-based tool-duplication deletion.

var csvfile3 string = `name,ring,quadrant,isNew,moved,description
Go,Trial,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Infrastructure,false,2,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing`

// TODO: Discuss handling ties.
//var csvfile4 string = `name,ring,quadrant,isNew,moved,description
//Python,Adopt,Language,false,0,Its a programming Language
//Visual Studio,Trial,Infrastructure,false,1,An IDE
//Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing`

var correctMergeMajority string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Infrastructure,false,2,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing
Python,Hold,Language,false,0,Its a programming Language`

var TestFiles []string = []string{"testFile1.csv", "testFile2.csv"}
var AllTestFiles []string = []string{"testFile1.csv", "testFile2.csv", "testFile3.csv"}

func createSomeCsv(count int) {

	if count > len(AllTestFiles) {
		log.Fatal("The count given is greater than the number of test files.")
	}
	var csvfile = csvfile1

	for i := 0; i < count; i++ {
		if i == 1 {
			csvfile = csvfile2
		} else if i == 2 {
			csvfile = csvfile3
		}

		err := os.WriteFile(AllTestFiles[i], []byte(csvfile), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createCsvFiles() {
	createSomeCsv(2)
}

func cleanSomeCsv(count int) {

	if count > len(AllTestFiles) {
		log.Fatal("The count given is greater than the number of test files.")
	}

	for i := 0; i < count; i++ {
		err := os.Remove(AllTestFiles[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanUp() {
	cleanSomeCsv(2)
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

func TestDuplicateRemoval(t *testing.T) {
	// Arrange clean-up after test finishes
	filename := "testFile0.csv"

	defer func() {
		err := os.Remove(filename)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()

	// Arrange input
	var csvfile0 string = "Go,Adopt,Language,true,0,Its a programming Language\n" +
		"Go,Adopt,Language,true,0,Its a programming Language\n" +
		"Python,Hold,Language,false,0,Its a programming Language\n"

	err := os.WriteFile(filename, []byte(csvfile0), 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Arrange csv file being closed, so it can be removed.
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()

	// Arrange expected output
	var expectedString string = "Go,Adopt,Language,true,0,Its a programming Language\n" +
		"Python,Hold,Language,false,0,Its a programming Language\n"

	// Arrange other variables to be used
	var set = make(map[string][]string)
	var buf bytes.Buffer

	// Act to call scanLine that calls duplicateRemoval() on each line
	scanFile(file, &buf, set)

	// Assert
	bufferString := buf.String()
	if bufferString != expectedString {
		t.Errorf("Buffer doesn't contain the expected data.\nExpected: %s \nActual: %s", expectedString, bufferString)
	}
}

func TestReadCsvData(t *testing.T) {
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
Visual Studio,Trial,Infrastructure,false,1,An IDE`

	bufferString := buf.String()
	if strings.Compare(bufferString, correctString) == 0 {
		t.Errorf("Buffer doesn't contain the correct data.\nExpected: %s\n Actual: %s", correctString, bufferString)
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
