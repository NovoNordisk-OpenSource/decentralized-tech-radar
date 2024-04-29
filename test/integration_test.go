package test

import (
	"os"
	"strings"
	"testing"

	html "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/HTML"
	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Merger"
	Reader "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
)

// Test Set up
var testFileName string = "ForTesting"

var csvTestString1 string = `name,ring,quadrant,isNew,moved,description
TestBlip1,Assess,Language,true,1,This is a description
TestBlip2,Adopt,Infrastructure,false,0,Also a description`

var csvTestString2 string = `name,ring,quadrant,isNew,moved,description
TestBlip3,Assess,Language,true,1,This is a description
TestBlip4,Adopt,Infrastructure,false,0,Also a description`

// Tests
// Integration test
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateCsvFile() {
	err := os.WriteFile(testFileName+"1.csv", []byte(csvTestString1), 0644)
	check(err)
	err = os.WriteFile(testFileName+"2.csv", []byte(csvTestString2), 0644)
	check(err)
}

func CleanUp() {
	os.Remove(testFileName + "1.csv")
	os.Remove(testFileName + "2.csv")
	os.Remove("index.html")
	os.Remove("Merged_file.csv")
	os.RemoveAll("cache/")
	os.Remove("specfile.txt")
	os.Remove("repos.txt")

	//Works on Unix and Windows
	os.Remove("tech_radar.exe")
	os.RemoveAll("cache")
}

func AssertIndexHTML(t *testing.T, html string) {
	//check if the index.html was created
	_, err := os.Stat("index.html")
	if os.IsNotExist(err) {
		t.Fatal("Expected HTML file was not created.")
	}

	//read content of the HTML file
	content, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("Could not read the generated HTML file: %v", err)
	}
	contentStr := string(content)

	//check if content contains expected string
	if !strings.Contains(contentStr, html) {
		t.Errorf("HTML doesn't contain the expected data\nContained:\n%s\nExpected:\n%s", contentStr, html)
	}
}

// Tests
// Integration test
func TestReaderAndWriter(t *testing.T) {
	// Set up
	CreateCsvFile()
	defer CleanUp()

	// Read test file
	specs := Reader.CsvToString(testFileName + "1.csv")
	html.GenerateHtml(specs)

	// Read index.html
	indexHTMLContent, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("Failed to read index.html: %v", err)
	}

	stringToCheck := `Factory("name,ring,quadrant,isNew,moved,description\nTestBlip1,Assess,Language,true,1,This is a description\nTestBlip2,Adopt,Infrastructure,false,0,Also a description").build();`

	if !strings.Contains(string(indexHTMLContent), stringToCheck) {
		t.Errorf("The content of HTML does not contain %s", stringToCheck)
	}
}

func TestMerger2Reader2Writer(t *testing.T) {
	// Set up
	CreateCsvFile()
	defer CleanUp()

	// Merge test csv files
	err := Merger.MergeCSV([]string{testFileName + "1.csv", testFileName + "2.csv"})
	if err != nil {
		t.Fatalf("MergeCSV() gave an error: %v", err)
	}

	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Merged file was not created")
	}

	// Read merged file
	specs := Reader.CsvToString("Merged_file.csv")

	// Generate html
	html.GenerateHtml(specs)

	// Read index.html
	indexHTMLContent, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("Failed to read index.html: %v", err)
	}
	
	stringToCheck := `Factory("name,ring,quadrant,isNew,moved,description\nTestBlip1,Assess,Language,true,1,This is a description\nTestBlip2,Adopt,Infrastructure,false,0,Also a description\nTestBlip3,Assess,Language,true,1,This is a description\nTestBlip4,Adopt,Infrastructure,false,0,Also a description\n").build();`

	if !strings.Contains(string(indexHTMLContent), stringToCheck) {
		t.Errorf("The content of HTML does not contain %s", stringToCheck)
	}
}
