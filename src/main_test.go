package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	Reader "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SpecReader"
	view "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/View"
)

// Test Set up
var testFileName0 string = "ForTesting"
var testFileName1 string = "ForTestingToo"

var csvTestString0 string = `name,ring,quadrant,isNew,moved,description
TestBlip1,Assess,Language,true,1,This is a description
TestBlip2,Adopt,Tool,false,0,Also a description`

var csvTestString1 string = `name,ring,quadrant,isNew,moved,description
TestBlip3,Hold,Tool,false,0,This too is a description
TestBlip4,Test,Language, true,1,Also a descriptive description`

var correctHTML string = `<html>
	<head>
		<title>Header 1</title>
	</head>
	<body>
		<h1 class="pageTitle">Header 1</h1>
		<ul>
			
					<li>Name: TestBlip1</li>
					<li>Quadrant: Language</li>
					<li>Ring: Assess</li>
					<li>Is new: true</li>
					<li>Moved: 1</li>
					<li>Desc: This is a description</li>
			
					<li>Name: TestBlip2</li>
					<li>Quadrant: Tool</li>
					<li>Ring: Adopt</li>
					<li>Is new: false</li>
					<li>Moved: 0</li>
					<li>Desc: Also a description</li>
			
		</ul>
	</body>
</html>`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createCsvFile(testFileNameType string, csvTestStringType string) {
	err := os.WriteFile(testFileNameType+".csv", []byte(csvTestStringType), 0644)
	check(err)
}

func cleanUp(testFileNameType string) {
	os.Remove(testFileNameType + ".csv")
	os.Remove("index.html")
	//Works on Unix and Windows
	os.Remove("tech_radar.exe")
}

func assertIndexHTML(t *testing.T) {
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

	//check if content contains exptected string
	if !strings.Contains(contentStr, correctHTML) {
		t.Errorf("HTML doesn't contain the expected data\nContained:\n%s", contentStr)
	}
}

// Tests
// Integration test
func TestReaderAndWriter(t *testing.T) {
	// Set up
	createCsvFile(testFileName0, csvTestString0)
	defer cleanUp(testFileName0)

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName0 + ".csv")
	view.GenerateHtml(specs)

	assertIndexHTML(t)
}

// End-to-end test
func TestEndToEnd(t *testing.T) {
	// Set up
	createCsvFile(testFileName0, csvTestString0)
	defer cleanUp(testFileName0)

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName0 + ".csv")
	view.GenerateHtml(specs)

	// Start program using CLI arguments
	os.Args = []string{"cmd", testFileName0 + ".csv"}
	//Works on Unix and Windows
	cmd := exec.Command("go", "build", "-o", "tech_radar.exe")
	cmd1 := exec.Command("./tech_radar.exe", "-file", testFileName0+".csv")

	_, err := cmd.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	cmd1_output, err := cmd1.Output()
	if err != nil {
		t.Fatalf("%v", err)
	} else if !strings.Contains(string(cmd1_output), "Opened csv file!") {
		t.Errorf("Output didn't match expected. %s", string(cmd1_output))
	}

	assertIndexHTML(t)
}
