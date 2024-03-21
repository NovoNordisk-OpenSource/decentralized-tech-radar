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

// TODO: [Nice to have] Automate the two vars below by using the two test-strings above.
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

var correctMergedHTML string = `<html>
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

					<li>Name: TestBlip3</li>
					<li>Quadrant: Hold</li>
					<li>Ring: Tools</li>
					<li>Is new: false</li>
					<li>Moved: 0</li>
					<li>Desc: This too is a description</li>

					<li>Name: TestBlip4</li>
					<li>Quadrant: Test</li>
					<li>Ring: Language</li>
					<li>Is new: true</li>
					<li>Moved: 1</li>
					<li>Desc: Also a descriptive description</li>
			
		</ul>
	</body>
</html>`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createCsvFile(amountOfTestFiles int) {
	err := os.WriteFile(testFileName0+".csv", []byte(csvTestString0), 0644)
	check(err)

	if amountOfTestFiles == 2 {
		err1 := os.WriteFile(testFileName1+".csv", []byte(csvTestString1), 0644)
		check(err1)
	}
}

func cleanUp(amountOfTestFiles int) {
	os.Remove(testFileName0 + ".csv")

	if amountOfTestFiles == 2 {
		os.Remove(testFileName1 + ".csv")
	}

	os.Remove("index.html")
	//Works on Unix and Windows
	os.Remove("tech_radar.exe")
}

func assertIndexHTML(t *testing.T, chosenCorrectHTML string) {
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
	if !strings.Contains(contentStr, chosenCorrectHTML) {
		t.Errorf("HTML doesn't contain the expected data\nContained:\n%s", contentStr)
	}
}

// Tests

// Unit-testing merger
func TestMerger(t *testing.T) {
	// Set up
	createCsvFile(2)
	defer cleanUp(2)

}

// Integration test
func TestReaderAndWriter(t *testing.T) {
	// Set up
	createCsvFile(1)
	defer cleanUp(1)

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName0 + ".csv")
	view.GenerateHtml(specs)

	assertIndexHTML(t, correctHTML)
}

// End-to-end test
func TestEndToEnd(t *testing.T) {
	// Set up
	createCsvFile(1)
	defer cleanUp(1)

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

	cmd1Output, err := cmd1.Output()
	if err != nil {
		t.Fatalf("%v", err)
	} else if !strings.Contains(string(cmd1Output), "Opened csv file!") {
		t.Errorf("Output didn't match expected. %s", string(cmd1Output))
	}

	assertIndexHTML(t, correctHTML)
}
