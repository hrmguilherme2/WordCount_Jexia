package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/chrislusf/glow/flow"
)

// CountWordRepeated is a representation of how much word we have in the text
var CountWordRepeated int

// LastWord is a representation the last word that we received to count
var LastWord string

// Result is a representation the final result e.g (word, 2)
var Result string

// BytesWrite is a representation of total Bytes Write in the file
var BytesWrite int

// BytesRead is a representation of total Bytes Readed in the file
var BytesRead int

// Regex purpose to remove all special characaters
var chars = []string{"]", "^", "!", "[", ".", "(", ")", "-", "?", `"`, `%`, `$`, `#`}

func main() {
	// Measure start execution time
	start := time.Now()
	// Parameters where InputFlag(-input) is the file that will be analyzed, and (-output) will be the result file
	InputFlag := flag.String("input", "metamorphosis.txt", "Path from the file that will be processed!")
	OutputFlag := flag.String("output", "Jexia.csv", "Path from the file that will be processed!")
	flag.Parse()
	parsingText(*InputFlag, *OutputFlag)
	// Measure end execution time
	ExecutionTime := time.Since(start)
	fmt.Println("Bytes Read:", BytesRead)
	fmt.Println("Bytes Write:", BytesWrite)
	fmt.Println("Execution Time: ", ExecutionTime)
}

// Customized errors message;
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// Function to remove any special char
func parsingString(word string, chars []string) string {
	CharsReceived := strings.Join(chars, "")
	RegexSlicer := regexp.MustCompile("[" + CharsReceived + "]")
	ProcessedString := RegexSlicer.ReplaceAllString(word, "")
	return strings.ToLower(ProcessedString)
}

// Function to manipulate the file and the result will be the word count in file
func parsingText(input string, output string) {
	if Exists(output) {
		os.Remove(output)
	} else {
		os.Create(output)
	}
	flow.New().TextFile(
		input, 7,
	).Partition(
		2,
	).Map(func(line string, ch chan string) {
		// println("Map:", line)
		for _, word := range strings.Split(line, " ") {
			ch <- parsingString(word, chars)
			// Bytes Read
			BytesRead += bytes.NewReader([]byte(word)).Len()
		}
	}).Filter(func(line string) bool {
		// println("Filter:", line)
		if line != "" && !strings.HasPrefix(line, " ") {
			return true
		}
		return false
	}).Sort(func(a string, b string) bool {
		if strings.Compare(a, b) < 0 {
			return true
		}
		return false
	}).Reduce(func(a string, b string) string {
		if LastWord == "" {
			LastWord = a
			CountWordRepeated = 1
		}
		if b == LastWord {
			CountWordRepeated++
		} else {
			Result = fmt.Sprintf("%s, %d", LastWord, CountWordRepeated)
			LastWord = b
			CountWordRepeated = 1
			// Write result in CSV
			FileCSV, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			checkError("Can't manage the file (Trying to open)...", err)
			WriteFile := csv.NewWriter(FileCSV)
			WriteFile.Write([]string{Result})
			WriteFile.Flush()
			// Bytes Write
			BytesWrite += bytes.NewReader([]byte(Result)).Len() * CountWordRepeated
		}
		return Result
	}).Run()

}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
