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
	// outputFile := os.Args[2]
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
	}
	converted := (string(fileContent))
	fileWords := strings.Fields(converted)
	fileWords = handleTags(fileWords)
	fmt.Println(fileWords)

}

func handleTags(fileWords []string) []string {
	for index, word:= range fileWords {
		switch word {
		case "(hex)":
			// convert the previous word from hex to decimal and replace it in the slice
			hexValue := fileWords[index-1]
			convertedValue, err := strconv.ParseInt(hexValue, 16, 64)
			if err != nil {
				fmt.Printf("Error decoding hex value: %v\n", err)
				continue
			}
			decimalValue := strconv.Itoa(int(convertedValue))
			fileWords[index-1] = decimalValue
			// remove the (hex) tag
			fileWords = append(fileWords[:index], fileWords[index+1:]...)
		}
		case "(up)":
			// Convert the previous word to Uppercase
			InputFileWords[i-1] = strings.ToUpper(InputFileWords[i-1])
			InputFileWords = append(InputFileWords[:i], InputFileWords[i+1:]...)
			i--
	}
	return fileWords
}