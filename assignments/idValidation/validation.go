package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validationCheck(social string) int {
	var sum int
	var numRegex = regexp.MustCompile(`^[0-9 ]+$`)

	if numRegex.MatchString(social) {
		social = strings.ReplaceAll(social, " ", "") // Remove spaces
		idArray := convertStringToArray(social)      // Get slice from function

		// Step 1: Process every second digit
		for i := len(idArray) - 2; i >= 0; i = i - 2 { // Fixed loop condition
			convertstrtoint, err := strconv.Atoi(idArray[i])
			if err != nil {
				panic(err)
			}
			doubledNum := convertstrtoint * 2

			if doubledNum > 9 {
				idArray[i] = strconv.Itoa(doubledNum - 9)
			} else {
				idArray[i] = strconv.Itoa(doubledNum)
			}
		}

		// Step 2: Sum up all digits
		for i := 0; i < len(idArray); i++ {
			convertedStr, err := strconv.Atoi(idArray[i])
			if err != nil {
				panic(err)
			}
			sum += convertedStr
		}
	}
	return sum
}

func convertStringToArray(str string) []string {
	array := make([]string, len(str)) // Correct slice allocation
	for index, value := range str {
		array[index] = string(value) // Proper assignment
	}
	return array
}

func main() {
	var socialNumber string = "4539 3195 0343 6467"
	validCheck := false
	result := validationCheck(socialNumber) // Test the function
	if result%10 == 0 {
		validCheck = true
	} else {
		validCheck = false
	}
	fmt.Println(validCheck)
}
