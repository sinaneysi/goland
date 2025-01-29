package main

import "fmt"

func main() {
	one := "GAGCCTACTAACGGGAT"
	two := "CATCGTAATGACGGCCT"
	arrayOne := make([]string, len(one))
	arrayTwo := make([]string, len(two))

	convertStringToArray(one, arrayOne)
	convertStringToArray(two, arrayTwo)

	result := hammingDistance(arrayOne, arrayTwo)
	fmt.Println(result)
}

func convertStringToArray(str string, array []string) []string {
	for index, value := range str {
		array[index] = string(value)
	}
	return array
}
func hammingDistance(firstArray []string, secondArray []string) int {
	mistakes := 0
	for i := 0; i < len(firstArray); i++ {
		if firstArray[i] != secondArray[i] {
			mistakes += 1
		}
	}
	return mistakes
}
