package HTML

import (
	"os"
	"strings"
	"testing"

	Reader "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
)

func TestGenerateHtml(t *testing.T) {
	//set up data
	blips := Reader.Blips{
		Blips: []Reader.Blip{
			{
				Name:        "Test name",
				Quadrant:    "Test quadrant",
				Ring:        "Test ring",
				IsNew:       true,
				Moved:       0,
				Description: "Test description",
			},
		},
	}

	//Generate the HTML file
	GenerateHtml(blips)

	//check if the index.html was created
	_, err := os.Stat(htmlFileName + ".html")
	if os.IsNotExist(err) {
		t.Fatal("Expected HTML was not created.")
	}

	//read content of the HTML file
	content, err := os.ReadFile(htmlFileName + ".html")
	if err != nil {
		t.Fatalf("Could not read the generated HTML file: %v", err)
	}
	
	//convert content to string and check if contains the test data
	contentStr := string(content)
	if !strings.Contains(contentStr, "Test name") ||
	!strings.Contains(contentStr, "Test quadrant") {
		t.Errorf("HTML doesn't contain the expected data.")
	}

	//clean up test after test is completed.
	err = os.Remove(htmlFileName + ".html")
	if err != nil {
		t.Fatalf("Failed to remove test HTML file: %v", err)
	}
}