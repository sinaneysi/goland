package main

import (
	"fmt"
	"regexp"
	"strings"
)

func check(social string) {

}
func main() {
	name := "this is sina"
	var socialNumber string = "3381381923"
	var numRegex = regexp.MustCompile(`^[0-9]+$`)
	if numRegex.MatchString(socialNumber) {

	}
	fmt.Println(strings.ReplaceAll(name, " ", ""))
}
