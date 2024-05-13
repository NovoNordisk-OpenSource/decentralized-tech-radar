package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func startProgram(t *testing.T) {
	// Uses CLI commands to start program
	// Works on Unix and Windows
	cmd := exec.Command("go", "build", "-o", "tech_radar.exe", "../src")
	_, err := cmd.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}
}

// End-to-end test
func TestEndToEnd(t *testing.T) {
	// Set up
	CreateCsvFile()
	defer CleanUp()

	// Read test file
	// specs := Reader.ReadCsvSpec(testFileName + "1.csv")
	// html.GenerateHtml(specs)
	//os.Args = []string{"cmd", testFileName + "1.csv"}

	startProgram(t)

	// Fetcher
	url := "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"
	data := []byte("examples/csv_templates/template.csv")
	os.WriteFile("./specfile.txt", data, 0644)
	specFile := "specfile.txt"

	cmd0 := exec.Command("./tech_radar.exe", "fetch", url, "main", specFile)
	_, err := cmd0.Output()
	if err != nil {
		t.Fatal(err)
	}

	// Merger
	// TODO: also merge ./cache/template.csv once it the spelling mistake in the header has been fixed on Novo
	cmd1 := exec.Command("./tech_radar.exe", "merge", "./"+testFileName+"1.csv", "./"+testFileName+"2.csv")
	_, err = cmd1.Output()
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat("Merged_file.csv")
	if os.IsNotExist(err) {
		t.Fatal("Failed to create Merged_file.csv")
	}

	// Generator
	cmd2 := exec.Command("./tech_radar.exe", "generate", "Merged_file.csv")
	_, err = cmd2.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Slices for assertion/checking with contains
	correctBlipNames := []string{
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

	// Check if index DOES NOT contain some of the Blips' names.
	for _, name := range correctBlipNames {
		if !strings.Contains(string(indexHTMLContent), name) {
			t.Errorf("Expected Blip-name %q not found in index.html", name)
		}
	}
}

func TestAddCmd(t *testing.T) {
	CreateCsvFile()
	defer CleanUp()

	startProgram(t)

	// Run the Add command
	cmd1 := exec.Command("./tech_radar.exe", "add", "ForTesting1.csv", "fakeLang", "assess", "language", "false", "0", "no lorem")
	_, err := cmd1.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Create slice with expected elements in the line appended
	correctAdd := []string{"fakeLang", "assess", "language", "false", "0", "no lorem"}

	// Read the ForTesting1.csv for asserts later
	readTestFile, err := os.ReadFile("ForTesting1.csv")
	if err != nil {
		t.Fatalf("Failed to read ForTesting1.csv: %v", err)
	}

	// Check ForTesting1 DOES NOT contain TestBlip1.
	if !strings.Contains(string(readTestFile), "TestBlip1") {
		t.Errorf("ForTesting1.csv does not contain blip-name: TestBlip1")
	}

	// Check if csv DOES NOT contain the fakeLang line.
	for _, lineElem := range correctAdd {
		if !strings.Contains(string(readTestFile), lineElem) {
			t.Errorf("Expected appended line %q not found in index.html", lineElem)
		}
	}
}

func TestE2EUsingFetcherFlags(t *testing.T) {
	// Set up
	os.Create("specfile.txt")
	err := os.WriteFile("specfile.txt", []byte("examples/csv_templates/template.csv"), 0644)
	if err != nil {
		t.Fatalf("Failed to create specfile.txt: %v", err)
	}
	os.Create("repos.txt")
	err = os.WriteFile("repos.txt", []byte("https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar main specfile.txt"), 0644)
	if err != nil {
		t.Fatalf("Failed to create repos.txt: %v", err)
	}

	CreateCsvFile()
	defer CleanUp()
	defer os.Remove("specfile.txt")

	// Works on Unix and Windows
	cmd := exec.Command("go", "build", "-o", "tech_radar.exe", "../src")
	_, err = cmd.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Fetch files using CLI arguments and flags
	cmd1 := exec.Command("./tech_radar.exe", "fetch", "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar", "--branch=main", "--whitelist=./specfile.txt")
	_, err = cmd1.Output()
	if err != nil {
		t.Fatal(err)
	}

	// Check if the file was downloaded
	_, err = os.Stat("cache/template.csv")
	if os.IsNotExist(err) {
		t.Fatal("Failed to create Merged_file.csv")
	}

	// Fetch files using file flag
	cmd2 := exec.Command("./tech_radar.exe", "fetch", "--repo-file=./repos.txt")
	_, err = cmd2.Output()
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("repos.txt", []byte("https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar"), 0644)
	if err != nil {
		t.Fatalf("Failed to create repos.txt: %v", err)
	}

	// Check combination of both flags
	cmd3 := exec.Command("./tech_radar.exe", "fetch", "--repo-file=./repos.txt", "--branch=main", "--whitelist=./specfile.txt")
	_, err = cmd3.Output()
	if err != nil {
		t.Fatal(err)
	}

}

func TestRemCmd(t *testing.T) {
	CreateCsvFile()
	defer CleanUp()

	// Works on Unix and Windows
	cmd := exec.Command("go", "build", "-o", "tech_radar.exe", "../src")
	_, err := cmd.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Run the Remove command
	cmd1 := exec.Command("./tech_radar.exe", "remove", "ForTesting1.csv", "TestBlip2", "Infrastructure")
	_, err = cmd1.Output()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Read the ForTesting1.csv for asserts later
	readTestFile, err := os.ReadFile("ForTesting1.csv")
	if err != nil {
		t.Fatalf("Failed to read ForTesting1.csv: %v", err)
	}

	// Check ForTesting1 still contains TestBlip1.
	if !strings.Contains(string(readTestFile), "TestBlip1") {
		t.Error("1: ForTesting1.csv does not contain blip-name: TestBlip1")
	}

	// Check if csv DOES NOT contain the removed line.
	if strings.Contains(string(readTestFile), "TestBlip2,Adopt,Infrastructure,false,0,Also a description") {
		t.Error("1: ForTesting1.csv contains TestBlip2 in Infrastructure")
	}
}
