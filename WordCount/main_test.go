package main

import (
	"io/ioutil"
	"testing"
)

func TestCreatingFileAndCheckTheResult(t *testing.T) {
	// Create the file
	parsingText("TestInput.txt", "TestOutput.csv")
	// Reading the file with results
	b, err := ioutil.ReadFile("TestOutput.csv")
	checkError("Cant open the file...", err)

	// Give the result from readed file to an variable
	actualResult := string(b)

	// What we expected
	var expectedResult = `"an, 1"
"file, 1"
"for, 1"
"is, 1"
"purpose, 1"
"test, 2"
`
	// Test result
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestRemovingSpecialChar(t *testing.T) {
	// Criteria chars to remove
	var chars = []string{"]", "^", "!", "[", ".", "(", ")", "-", "?", `"`, `%`, `$`, `#`}
	// What we Received
	actualResult := parsingString("Jexia,!", chars)
	// What we expected
	var expectedResult = "jexia"
	// Test result
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
