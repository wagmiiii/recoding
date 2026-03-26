package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	converted := (string(fileContent))
	fileWords := strings.Fields(converted)
	fileWords = handleTags(fileWords)

	outputString := strings.Join(fileWords, " ")
	err = os.WriteFile(outputFile, []byte(outputString), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
	}
}

func handleTags(fileWords []string) []string {
	for i := 0; i < len(fileWords); i++ {
		switch fileWords[i] {
		case "(hex)":
			if i > 0 {
				// convert the previous word from hex to decimal
				val, err := strconv.ParseInt(fileWords[i-1], 16, 64)
				if err == nil {
					fileWords[i-1] = strconv.FormatInt(val, 10)
				}
				// remove the (hex) tag
				fileWords = append(fileWords[:i], fileWords[i+1:]...)
				i-- // adjust index after removal
			}
		case "(bin)":
			if i > 0 {
				// convert the previous word from binary to decimal
				val, err := strconv.ParseInt(fileWords[i-1], 2, 64)
				if err == nil {
					fileWords[i-1] = strconv.FormatInt(val, 10)
				}
				// remove the (bin) tag
				fileWords = append(fileWords[:i], fileWords[i+1:]...)
				i--
			}
		case "(up)":
			if i > 0 {
				// Convert the previous word to Uppercase
				fileWords[i-1] = strings.ToUpper(fileWords[i-1])
				// remove the (up) tag
				fileWords = append(fileWords[:i], fileWords[i+1:]...)
				i--
			}
		}
	}
	return fileWords
}
