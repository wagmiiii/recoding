package main

import (
	"fmt"
	"log"
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
		fmt.Println("Error reading input file:", err)
		return
	}

	converted := (string(fileContent))
	fileWords := strings.Fields(converted)

	fileWords = handleTags(fileWords)

	fmt.Println(fileWords)
}

func handleTags(fileWords []string) []string {
	for index := 0; index < len(fileWords); index++ {
		word := fileWords[index]
		if index == 0 && (word == "(hex)" || word == "(bin)" || word == "(up)" || word == "(low)" || word == "(cap)") {
			log.Fatal("\nError: tag cannot be the first word in the file.")
		}
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
			index-- // Adjust index after removing the tag

		case "(bin)":
			//convert the previous word from bin to decimal
			binValue := fileWords[index-1]

			convertedValue, err := strconv.ParseInt(binValue, 2, 64)
			if err != nil {
				fmt.Println("Error decoding bin value ", err)
				continue
			}

			decimalValue := strconv.Itoa(int(convertedValue))
			fileWords[index-1] = decimalValue

			fileWords = append(fileWords[:index], fileWords[index+1:]...)
			index-- // Adjust index after removing the tag
			
		case "(up)" :
			fileWords[index-1] = strings.ToUpper(fileWords[index-1])
			fileWords = append(fileWords[:index], fileWords[index+1:]...)
			index-- // Adjust index after removing the tag

		case "(low)" :
			fileWords[index-1] = strings.ToLower(fileWords[index-1])
			fileWords = append(fileWords[:index], fileWords[index+1:]...) 
			index-- // Adjust index after removing the tag

		case "(cap)" :
			fileWords[index-1] = capFirstLetter(fileWords[index-1])
			fileWords = append(fileWords[:index], fileWords[index+1:]...)
			index-- // Adjust index after removing the tag
		}
		
	}

	return fileWords
}

func capFirstLetter(word string) string {
	word = strings.ToLower(word) // handle words with all caps by first turning to lowercase
	firstLetter := strings.ToUpper(string(word[0])) // /
	restLetters := word[1:]
	return firstLetter + restLetters
}
