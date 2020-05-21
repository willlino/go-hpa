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

	if string(out) != "William Rocks!" {
		t.Errorf("Expected %s, got %s", "William Rocks!", out)
	}
}

func TestDoSqrt(t *testing.T) {
	var input float64 = 9

	result := doSqrt(input)

	if result != 3 {
		t.Errorf("Value 3 expected, but got: %v", result)
	}
}

func TestValidateQueryString_ShouldNotBeEmpty(t *testing.T) {
	input := ""
	messageExpected := "The input 'value' cannot be empty"

	errors := validateQueryString(input)

	if errors[0] != messageExpected {
		t.Errorf("The list of erros should contain: %v", messageExpected)
	}
}

func TestValidateQueryString_IsValid(t *testing.T) {
	input := "10"

	errors := validateQueryString(input)

	if len(errors) > 0 {
		t.Errorf("The list of erros should not contain any error")
	}
}

func TestValidateQueryString_AddErrorIsNotNumber(t *testing.T) {
	input := "stringTest"
	messageExpected := "The input 'value' must be a float"

	errors := validateQueryString(input)

	if errors[0] != messageExpected {
		t.Errorf("The list of erros should contain: %v", messageExpected)
	}
}
