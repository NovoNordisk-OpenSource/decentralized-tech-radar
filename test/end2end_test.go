package test

/*import (
	"os"
	"os/exec"
	"strings"
	"testing"

	html "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/HTML"
	Reader "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
)

// End-to-end test
func TestEndToEnd(t *testing.T) {
	// Set up
	CreateCsvFile()
	defer CleanUp()

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName + "1.csv")
	html.GenerateHtml(specs)

	// Start program using CLI arguments
	os.Args = []string{"cmd", testFileName + "1.csv"}
	//Works on Unix and Windows
	cmd := exec.Command("go", "build", "-o", "tech_radar.exe", "../src")
	_, err := cmd.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	cmd1 := exec.Command("./tech_radar.exe", "-merge", testFileName+"1.csv"+" "+testFileName+"2.csv")
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
			<li>Quadrant: Infrastructure</li>
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
			<li>Quadrant: Infrastructure</li>
			<li>Ring: Adopt</li>
			<li>Is new: false</li>
			<li>Moved: 0</li>
			<li>Desc: Also a description</li>

		</ul>
	</body>
</html>`
	AssertIndexHTML(t, correctHTML)
}*/
