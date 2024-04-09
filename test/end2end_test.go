package test

import (
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

	// Fetcher
	url := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	data := []byte("examples/csv_templates/template.csv")
	os.WriteFile("./specfile.txt", data, 0644)
	specFile := "specfile.txt"

	cmd0 := exec.Command("./tech_radar.exe", "fetch", url, "main", specFile)
	_, err = cmd0.Output()
	if err != nil {
		t.Fatal(err)
	}

	cmd1 := exec.Command("./tech_radar.exe", "merge", "./cache/template.csv", "./cache/"+testFileName+"1.csv", "./cache/"+testFileName+"2.csv")
	_, err = cmd1.Output()
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Failed to create Merged_file.csv")
	}

	cmd2 := exec.Command("./tech_radar.exe", "generate", "Merged_file.csv")
	cmd2_output, err := cmd2.Output()
	if err != nil {
		t.Fatalf("%v", err)
	} else if !strings.Contains(string(cmd2_output), "Opened csv file!") {
		t.Errorf("Output didn't match expected. %s", string(cmd2_output))
	}

	correctBlipNames := []string{
		"Python",
		"web",
		"react",
		"TestBlip1",
		"TestBlip2",
		"TestBlip3",
		"TestBlip4",
	}

	// Read index.html
	indexHTMLContent, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("Failed to read index.html: %v", err)
	}

	// Check if index contains some of the Blips' names.
	for _, name := range correctBlipNames {
		if !strings.Contains(string(indexHTMLContent), name) {
			t.Errorf("Expected Blip-name %q not found in index.html", name)
		}
	}

}
