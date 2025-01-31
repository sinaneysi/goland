package hamingFunc

// ConvertStringToArray converts a string into a slice of strings
func convertStringToArray(str string, array []string) []string {
	for index, value := range str {
		array[index] = string(value)
	}
	return array
}

// HammingDistance calculates the Hamming distance between two slices
func HammingDistance(firstArray, secondArray []string) int {
	mistakes := 0
	for i := 0; i < len(firstArray); i++ {
		if firstArray[i] != secondArray[i] {
			mistakes++
		}
	}
	return mistakes
}
