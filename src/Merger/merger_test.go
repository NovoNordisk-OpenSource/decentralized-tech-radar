package Merger

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"go.uber.org/zap"
)

var csvfile1 string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Languages & frameworks,true,0,Its a programming Language
Visual Studio Code,Trial,Platforms,false,2,An IDE
Dagger IO,Assess,Tools,true,1,Its a workflow thing`

var csvfile2 string = `name,ring,quadrant,isNew,moved,description
Python,Hold,Languages & frameworks,false,0,Its a programming Language
Visual Studio,Trial,Platforms,false,1,An IDE
Dagger IO,Assess,Tools,true,1,Its a workflow thing`

var correctMerge string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Languages & frameworks,true,0,Its a programming Language
Visual Studio Code,Trial,Platforms,false,2,An IDE
Visual Studio,Trial,Platforms,false,1,An IDE
Dagger IO,Assess,Tools,true,1,Its a workflow thing
Python,Hold,Languages & frameworks,false,0,Its a programming Language`

// The strings below are used exclusively to test
// ReadCsvFile's majority-vote-based tool-duplication deletion.

var csvfile3 string = `name,ring,quadrant,isNew,moved,description
Go,Trial,Languages & frameworks,true,0,Its a programming Language
Visual Studio Code,Trial,Platforms,false,2,An IDE
Dagger IO,Assess,Tools,true,1,Its a workflow thing`

// TODO: Discuss handling ties.
//var csvfile4 string = `name,ring,quadrant,isNew,moved,description
//Python,Adopt,Language,false,0,Its a programming Language
//Visual Studio,Trial,Infrastructure,false,1,An IDE
//Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing`

var correctMergeMajority string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Languages & frameworks,true,0,Its a programming Language
Visual Studio Code,Trial,Platforms,false,2,An IDE
Dagger IO,Assess,Tools,true,1,Its a workflow thing
Python,Hold,Languages & frameworks,false,0,Its a programming Language`

var oldTestFiles []string = []string{"testFile1.csv", "testFile2.csv"}
var testFiles []string = []string{"testFile1.csv", "testFile2.csv", "testFile3.csv"}

func createSomeCsv(count int) {

	if count > len(testFiles) {
		log.Fatal("The count given is greater than the number of test files.")
	}
	var csvfile = csvfile1

	for i := 0; i < count; i++ {
		if i == 1 {
			csvfile = csvfile2
		} else if i == 2 {
			csvfile = csvfile3
		}

		err := os.WriteFile(testFiles[i], []byte(csvfile), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanSomeCsv(count int) {

	if count > len(testFiles) {
		log.Fatal("The count given is greater than the number of test files.")
	}

	for i := 0; i < count; i++ {
		err := os.Remove(testFiles[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanUp(count int) {
	cleanSomeCsv(count)
	os.Remove(("Merge_log.txt"))
	os.Remove("Merged_file.csv")
	os.RemoveAll("./cache")
}

func TestGetHeader(t *testing.T) {
	createSomeCsv(1)
	defer cleanUp(1)
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
	var csvfile0 string = "Go,Adopt,Languages & frameworks,true,0,Its a programming Language\n" +
		"Go,Adopt,Languages & frameworks,true,0,Its a programming Language\n" +
		"Python,Hold,Languages & frameworks,false,0,Its a programming Language\n"

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
	var expectedString string = "Go,Adopt,Languages & frameworks,true,0,Its a programming Language\n" +
		"Python,Hold,Languages & frameworks,false,0,Its a programming Language\n"

	// Arrange other variables to be used
	var set = make(map[string][]string)
	blips := make(map[string]map[string]byte)
	
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// Act to call scanFile that calls duplicateRemoval() on each line
	Fcfs{}.scanFile(file, set, sugar)

	// Assert
	for line := range blips {
		if !(strings.Contains(expectedString, line)) {
			t.Errorf("The line from the blips is not in the expected string\nExpected string: %s\nBlips line: %s", expectedString, line)
		}
	}
}

func TestReadCsvData(t *testing.T) {
	createSomeCsv(2)
	defer cleanUp(2)

	var buf bytes.Buffer
	Fcfs{}.MergeFiles(&buf, "./testFile1.csv", "./testFile2.csv")

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

	correctString := `Go,Adopt,Languages & frameworks,true,0,Its a programming Language
Visual Studio Code,Trial,Platforms,false,2,An IDE
Dagger IO,Assess,Tools,true,1,Its a workflow thing
Python,Hold,Languages & frameworks,false,0,Its a programming Language
Visual Studio,Trial,Platforms,false,1,An IDE`

	bufferString := buf.String()
	if strings.Compare(bufferString, correctString) == 0 {
		t.Errorf("Buffer doesn't contain the correct data.\nExpected: %s\n Actual: %s", correctString, bufferString)
	}
}

func TestMergeCSV(t *testing.T) {
	// Setup
	createSomeCsv(2)
	defer cleanUp(2)

	mergeTestFiles := append([]string{}, testFiles[0], testFiles[1])

	// Call function
	err := MergeCSV(mergeTestFiles, Fcfs{})
	if err != nil {
		t.Fatalf("MergeCSV gave an error: %v", err)
	}

	// Check that file exists
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Merged_file.csv was not found")
	}

	// Check that file is merged correctly
	contentFile, err := os.Open("Merged_file.csv")
	if err != nil {
		t.Fatalf("Merged_file.csv could not be opened: %v", err)
	}
	defer contentFile.Close()

	scanner := bufio.NewScanner(contentFile)
  
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(correctMerge, line) {
			t.Errorf("Merged file doesn't contain the expected data\nExpected: %s\nContained: %s", correctMerge, line)
		}
	}
}

func TestMergeFromFolder(t *testing.T) {
	createSomeCsv(2)

	// Note: It will print "remove testFileX.csv: ..."
	// This is because the files have been renamed/moved into cache
	// So it cannot find the original file names
	// But cleanup also handles removing the cache-folder, so this is ok.

	defer cleanUp(2)

	err := os.Mkdir("cache", 0700)
	if err != nil {
		t.Fatal(err)
	}
	os.Rename(testFiles[0], "./cache/"+testFiles[0])
	os.Rename(testFiles[1], "./cache/"+testFiles[1])

	err = MergeFromFolder("./cache", Fcfs{})
	if err != nil {
		t.Fatal(err)
	}

	// Check that file exists
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Merged_file.csv was not found")
	}

	// Check that file is merged correctly
	contentFile, err := os.Open("Merged_file.csv")
	if err != nil {
		t.Fatalf("Merged_file.csv could not be opened: %v", err)
	}
	defer contentFile.Close()

	scanner := bufio.NewScanner(contentFile)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(correctMerge, line) {
			t.Errorf("Merged file doesn't contain the expected data\nExpected: %s\nContained: %s", correctMerge, line)
		}
	}
}
