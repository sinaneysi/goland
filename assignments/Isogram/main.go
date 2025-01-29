package main

import (
	"fmt"
	"strings"
	"unicode"
)

func trimNonAlphabetic(input string) string {
	var result strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) { // Keep only letters (A-Z, a-z)
			result.WriteRune(char)
		}
	}
	return result.String()
}
func isogram(str string, array []string) bool {
	var ifIsogram bool = true
	for index, value := range str {
		array[index] = string(value)
	}
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if array[i] == array[j] {
				ifIsogram = false
				break
			}
		}
	}
	return ifIsogram
}
func main() {
	var name string
	fmt.Println("Enter your name to check if its isogram or not: ")
	fmt.Scanln(&name)
	name = trimNonAlphabetic(name)
	nameArray := make([]string, len(name))
	ifIsogram := isogram(name, nameArray)
	fmt.Println(ifIsogram)

}
