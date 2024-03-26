package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	Reader "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SpecReader"
	view "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/View"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/Merger"
)

// Test Set up
var testFileName string = "ForTesting"

var csvTestString1 string = `name,ring,quadrant,isNew,moved,description
TestBlip1,Assess,Language,true,1,This is a description
TestBlip2,Adopt,Tool,false,0,Also a description`

var csvTestString2 string = `name,ring,quadrant,isNew,moved,description
TestBlip3,Assess,Language,true,1,This is a description
TestBlip4,Adopt,Tool,false,0,Also a description`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createCsvFile() {
	err := os.WriteFile(testFileName + "1.csv", []byte(csvTestString1), 0644)
	check(err)
	err = os.WriteFile(testFileName + "2.csv", []byte(csvTestString2), 0644)
	check(err)
}

func cleanUp() {
	os.Remove(testFileName + "1.csv")
	os.Remove(testFileName + "2.csv")
	os.Remove("index.html")
	os.Remove("Merged_file.csv")
	//Works on Unix and Windows
	os.Remove("tech_radar.exe")
}

func assertIndexHTML(t *testing.T, html string) {
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
	if !strings.Contains(contentStr, html) {
		t.Errorf("HTML doesn't contain the expected data\nContained:\n%s\nExpected:\n%s", contentStr, html)
	}
}

// Tests
// Integration test
func TestReaderAndWriter(t *testing.T) {
	// Set up
	createCsvFile()
	defer cleanUp()

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName + "1.csv")
	view.GenerateHtml(specs)
	
	correctHTML := `<html>
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
	assertIndexHTML(t, correctHTML)
}

func TestMerger2Reader2Writer(t *testing.T) {
	// Set up
	createCsvFile()
	defer cleanUp()

	// Merge test csv files
	Merger.MergeCSV([]string{testFileName + "1.csv", testFileName + "2.csv"})
	_, err := os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Merged file was not created")
	}

	// Read merged file
	specs := Reader.ReadCsvSpec("Merged_file.csv")

	// Generate html
	view.GenerateHtml(specs)

	correctHTML := `<html>
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
			<li>Quadrant: Language</li>
			<li>Ring: Assess</li>
			<li>Is new: true</li>
			<li>Moved: 1</li>
			<li>Desc: This is a description</li>
			
			<li>Name: TestBlip4</li>
			<li>Quadrant: Tool</li>
			<li>Ring: Adopt</li>
			<li>Is new: false</li>
			<li>Moved: 0</li>
			<li>Desc: Also a description</li>
			
		</ul>
	</body>
</html>`
	assertIndexHTML(t, correctHTML)
}

// End-to-end test
func TestEndToEnd(t *testing.T) {
	// Set up
	createCsvFile()
	defer cleanUp()

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName + "1.csv")
	view.GenerateHtml(specs)

	// Start program using CLI arguments
	os.Args = []string{"cmd", testFileName + "1.csv"}
	//Works on Unix and Windows
	cmd := exec.Command("go", "build", "-o", "tech_radar.exe")
	_, err := cmd.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	cmd1 := exec.Command("./tech_radar.exe", "-merge", testFileName + "1.csv" + " " + testFileName + "2.csv")
	_, err = cmd1.Output()
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Failed to create Merged_file.csv")
	}

	cmd2 := exec.Command("./tech_radar.exe", "-file", "Merged_file.csv")
	cmd2_output, err := cmd2.Output()
	if err != nil {
		t.Fatalf("%v", err)
	} else if !strings.Contains(string(cmd2_output), "Opened csv file!") {
		t.Errorf("Output didn't match expected. %s", string(cmd2_output))
	}

	correctHTML := `<html>
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
			<li>Quadrant: Language</li>
			<li>Ring: Assess</li>
			<li>Is new: true</li>
			<li>Moved: 1</li>
			<li>Desc: This is a description</li>
			
			<li>Name: TestBlip4</li>
			<li>Quadrant: Tool</li>
			<li>Ring: Adopt</li>
			<li>Is new: false</li>
			<li>Moved: 0</li>
			<li>Desc: Also a description</li>
			
		</ul>
	</body>
</html>`
	assertIndexHTML(t, correctHTML)
}