package main

import "fmt"

func main() {
	for {
		var name string
		fmt.Println("Please enter a name:")
		fmt.Scanln(&name)
		if name == "" {
			fmt.Println("One for you, one for me.")
		} else {
			fmt.Printf("One for %s, one for me.\n", name)
		}
	}
}
