package Verifier

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

var csvfile1 string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Infrastructure,false,2,An IDE
Dagger IO,Assess,Infrastructure,true,1,Its a workflow thing`

func createCsvFiles(csvfile string) {
	err := os.WriteFile("testFile1.csv", []byte(csvfile), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("testFile2.csv", []byte(csvfile), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanUp() {
	os.Remove("testFile1.csv")
	os.Remove("testFile2.csv")
	os.Remove("tempfile.csv")
}

func TestVerifierFunctionDuplicateDeletion(t *testing.T) {
	createCsvFiles(csvfile1)
	defer cleanUp()

	Verifier("./testFile1.csv", "./testFile2.csv")

	csv1, err := os.ReadFile("./testFile1.csv")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(csv1), csvfile1) {
		t.Fatalf("csvFile1 does not match expected output.\nExpected: %s \n Actual: %s", csvfile1, csv1)
	}

	csv2, err := os.ReadFile("./testFile2.csv")
	if err != nil {
		t.Fatal(err)
	}

	if string(csv2) != "name,ring,quadrant,isNew,moved,description\n" {
		t.Fatalf("csvFile2 does not match expected output.\nExpected: name,ring,quadrant,isNew,moved,description \nActual: %s",csv2)
	}
}

var csvfile2 string = `name,ring,quadrant,isNew,moved,description
Go;Adopt?Language:true:0_Its a programming Language
Visual Studio Code:Trial^Tool:false;2:An IDE
Dagger IO;Assess*Tool+true?1_Its a workflow thing`

func TestCSVWrongFormatError(t *testing.T) {
	createCsvFiles(csvfile2)
	defer cleanUp()

	err := Verifier("./testFile1.csv")

	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestCheckHeaderCorrectHeader(t *testing.T) {
	createCsvFiles(csvfile1)
	defer cleanUp()

	file, err := os.Open("./testFile1.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	value := checkHeader(scanner.Text())
	if !value {
		t.Errorf("checkHeader returned the wrong value\n\tGot: %t\n\tExpected: %t", value, true)
	}
}

func TestCheckDatalineCorrectData(t *testing.T) {
	createCsvFiles(csvfile1)
	defer cleanUp()

	file, err := os.Open("./testFile1.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip header
	scanner.Scan()

	value := checkDataLine(scanner.Text())
	if !value {
		t.Errorf("checkDataline returned the wrong value\n\tGot: %t\n\tExpected: %t", value, true)
	}
}