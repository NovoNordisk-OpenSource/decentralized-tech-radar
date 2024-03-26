package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	html "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/HTML"
	Reader "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
)

// End-to-end test
func TestEndToEnd(t *testing.T) {
	// Set up
	createCsvFile()
	defer cleanUp()

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName + ".csv")
	html.GenerateHtml(specs)

	// Start program using CLI arguments
	os.Args = []string{"cmd", testFileName + ".csv"}
	//Works on Unix and Windows
	cmd := exec.Command("go", "build", "-C", "../src", "-o", "tech_radar.exe")
	cmd1 := exec.Command("../src/tech_radar.exe", "-file", testFileName+".csv")

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
