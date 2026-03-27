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
		switch word {

		case "(hex)":
			if index == 0 {
				fmt.Println("Error: (hex) tag cannot be the first word in the file.")
				return nil
			}
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
			if index == 0 {
				fmt.Println("Error: (bin) tag cannot be the first word in the file.")
				continue
			}
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
			if index == 0 {
				fmt.Println("Error: (up) tag cannot be the first word in the file.")
				continue
			}

			fileWords[index-1] = strings.ToUpper(fileWords[index-1])
			fileWords = append(fileWords[:index], fileWords[index+1:]...)
			index-- // Adjust index after removing the tag

		case "(low)" :
			if index == 0 {
				fmt.Println("Error: (low) tag cannot be the first word in the file.")
				continue
			}

			fileWords[index-1] = strings.ToLower(fileWords[index-1])
			fileWords = append(fileWords[:index], fileWords[index+1:]...) 
			index-- // Adjust index after removing the tag
		}
		
	}

	return fileWords
}
