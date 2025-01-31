package main

import (
	"assignment/hamingDistance/hamingFunc"
	"fmt"
)

func main() {
	one := "GAGCCTACTAACGGGAT"
	two := "CATCGTAATGACGGCCT"
	arrayOne := make([]string, len(one))
	arrayTwo := make([]string, len(two))
	
	hamingFunc.ConvertStringToArray(one, arrayOne)
	hamingFunc.ConvertStringToArray(two, arrayTwo)

	result := hamingFunc.HammingDistance(arrayOne, arrayTwo)
	fmt.Println(result)
}
