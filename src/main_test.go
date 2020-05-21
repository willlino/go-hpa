package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printMessage()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "Code.education Rocks!" {
		t.Errorf("Expected %s, got %s", "Code.education Rocks!", out)
	}
}

func TestDoSqrt(t *testing.T) {
	var input float64 = 9

	result := doSqrt(input)

	if result != 3 {
		t.Errorf("Value 3 expected, but got: %v", result)
	}
}
