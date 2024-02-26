package test

import (
	"testing"
)

func TestMain(t *testing.T) {
	output := SayHello("mate")
	if output != "Hello, mate" {
		t.Error("Output doesn't match")
	}
}
