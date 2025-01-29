package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int = 105
	result := calculateDiviseablity(number)
	fmt.Println(result)
}
func calculateDiviseablity(num int) string {
	result := ""
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(num)
	}
	return result
}
