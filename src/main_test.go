package main

import (
	"testing"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SayHello"
)

// unit test
func TestSayHello(t *testing.T) {
	output := SayHello.SayHello("mate")
	if output != "Hello, mate" {
		t.Error("Output doesn't match")
	}
}

// E2E test
/*func TestMain(t *testing.T) {
	name := "Naanaualjoth"
	cmd := exec.Command("cd", "../dist", "&&", "./main", "-name", name)

	cmd_output, err := cmd.Output()
	if err != nil {
		t.Errorf("%v", err)
	} else if string(cmd_output) != (fmt.Sprintf("Hello, %s", name)) {
		t.Errorf("Output didn't match expected. %s", string(cmd_output))
	}
}*/
