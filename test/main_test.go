package test

import (
	"testing"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SayHello"
)

func TestSayHello(t *testing.T) {
	output := SayHello.SayHello("mate")
	if output != "Hello, mate" {
		t.Error("Output doesn't match")
	}
}


