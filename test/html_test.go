package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	view "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/View"
)

/**
 * TestHtmlWriting
 * creates the index.html and tests the output of the html file if it exists
 * simply it checks if the string greeting exist in the html file
 */

func TestHtmlWriting(t *testing.T) {
	//sets up the greeting and index.html file
	greeting := "This is a test for html writing"
	view.MakeHtml(greeting)

	//read file
	content, err := os.ReadFile("index.html")
	if err != nil {
		panic(err)
	}

	strContent := string(content)

	// sees if the output contains the greeting
	if !strings.Contains(strContent, greeting) {
		t.Error("Output doesn't match")
	}

	//remove the index.html file after use
	if err := os.Remove("index.html"); err != nil {
		fmt.Print(err)
	}
}
